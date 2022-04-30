package chschema

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/uptrace/go-clickhouse/ch/chproto"
)

type ArrayColumnar interface {
	WriteOffset(wr *chproto.Writer, offset int) int
	WriteData(wr *chproto.Writer) error
}

//------------------------------------------------------------------------------

type ArrayColumnOf[T any] struct {
	Column [][]T
	elem   Columnar
}

func (c *ArrayColumnOf[T]) Reset(numRow int) {
	if cap(c.Column) >= numRow {
		c.Column = c.Column[:0]
	} else {
		c.Column = make([][]T, 0, numRow)
	}
}

func (c *ArrayColumnOf[T]) Set(v any) {
	c.Column = v.([][]T)
}

func (c *ArrayColumnOf[T]) Value() any {
	return c.Column
}

func (c *ArrayColumnOf[T]) Nullable(nulls UInt8Column) any {
	panic("not implemented")
}

func (c *ArrayColumnOf[T]) Len() int {
	return len(c.Column)
}

func (c *ArrayColumnOf[T]) Index(idx int) any {
	return c.Column[idx]
}

func (c *ArrayColumnOf[T]) Slice(s, e int) any {
	return c.Column[s:e]
}

func (c *ArrayColumnOf[T]) ConvertAssign(idx int, v reflect.Value) error {
	v.Set(reflect.ValueOf(c.Column[idx]))
	return nil
}

func (c *ArrayColumnOf[T]) AppendValue(v reflect.Value) {
	ptr := unsafe.Pointer(v.UnsafeAddr())
	c.AppendPointer(v.Type(), ptr)
}

func (c *ArrayColumnOf[T]) AppendPointer(typ reflect.Type, ptr unsafe.Pointer) {
	c.Column = append(c.Column, *(*[]T)(ptr))
}

func (c *ArrayColumnOf[T]) ReadFrom(rd *chproto.Reader, numRow int) error {
	if cap(c.Column) >= numRow {
		c.Column = c.Column[:numRow]
	} else {
		c.Column = make([][]T, numRow)
	}

	if numRow == 0 {
		return nil
	}

	offsets := make([]int, numRow)
	for i := 0; i < numRow; i++ {
		offset, err := rd.UInt64()
		if err != nil {
			return err
		}
		offsets[i] = int(offset)
	}

	if err := c.elem.ReadFrom(rd, offsets[len(offsets)-1]); err != nil {
		return err
	}

	var prev int
	for i, offset := range offsets {
		c.Column[i] = c.elem.Slice(prev, offset).([]T)
		prev = offset
	}

	return nil
}

func (c *ArrayColumnOf[T]) WriteTo(wr *chproto.Writer) error {
	_ = c.WriteOffset(wr, 0)
	return c.WriteData(wr)
}

var _ ArrayColumnar = (*Int64ArrayColumn)(nil)

func (c *ArrayColumnOf[T]) WriteOffset(wr *chproto.Writer, offset int) int {
	for _, el := range c.Column {
		offset += len(el)
		wr.UInt64(uint64(offset))
	}
	return offset
}

func (c *ArrayColumnOf[T]) WriteData(wr *chproto.Writer) error {
	for _, ss := range c.Column {
		c.elem.Set(ss)
		if err := c.elem.WriteTo(wr); err != nil {
			return err
		}
	}
	return nil
}

//------------------------------------------------------------------------------

type Int64ArrayColumn struct {
	ArrayColumnOf[int64]
}

var _ Columnar = (*Int64ArrayColumn)(nil)

func NewInt64ArrayColumn(typ reflect.Type, chType string, numRow int) Columnar {
	return &Int64ArrayColumn{
		ArrayColumnOf: ArrayColumnOf[int64]{
			Column: make([][]int64, 0, numRow),
			elem:   NewInt64Column(typ.Elem(), "", 0),
		},
	}
}

func (c *Int64ArrayColumn) Type() reflect.Type {
	return int64SliceType
}

//------------------------------------------------------------------------------

type Uint64ArrayColumn struct {
	ArrayColumnOf[uint64]
}

var _ Columnar = (*Uint64ArrayColumn)(nil)

func NewUint64ArrayColumn(typ reflect.Type, chType string, numRow int) Columnar {
	return &Uint64ArrayColumn{
		ArrayColumnOf: ArrayColumnOf[uint64]{
			Column: make([][]uint64, 0, numRow),
			elem:   NewUInt64Column(typ.Elem(), "", 0),
		},
	}
}

func (c *Uint64ArrayColumn) Type() reflect.Type {
	return uint64SliceType
}

//------------------------------------------------------------------------------

type Float64ArrayColumn struct {
	ArrayColumnOf[float64]
}

var _ Columnar = (*Float64ArrayColumn)(nil)

func NewFloat64ArrayColumn(typ reflect.Type, chType string, numRow int) Columnar {
	return &Float64ArrayColumn{
		ArrayColumnOf: ArrayColumnOf[float64]{
			Column: make([][]float64, 0, numRow),
			elem:   NewFloat64Column(typ.Elem(), "", 0),
		},
	}
}

func (c *Float64ArrayColumn) Type() reflect.Type {
	return float64SliceType
}

//------------------------------------------------------------------------------

type StringArrayColumn struct {
	Column     [][]string
	elem       Columnar
	stringElem *StringColumn
	lcElem     *LCStringColumn
}

var _ Columnar = (*StringArrayColumn)(nil)

func NewStringArrayColumn(typ reflect.Type, chType string, numRow int) Columnar {
	if _, funcType := aggFuncNameAndType(chType); funcType != "" {
		chType = funcType
	}
	elemType := chArrayElemType(chType)
	if elemType == "" {
		panic(fmt.Errorf("invalid array type: %q (Go type is %s)",
			chType, typ.String()))
	}

	columnar := NewColumn(typ.Elem(), elemType, 0)
	var stringElem *StringColumn
	var lcElem *LCStringColumn

	switch v := columnar.(type) {
	case *StringColumn:
		stringElem = v
	case *LCStringColumn:
		stringElem = &v.StringColumn
		lcElem = v
		columnar = &ArrayLCStringColumn{v}
	case *EnumColumn:
		stringElem = &v.StringColumn
	default:
		panic(fmt.Errorf("unsupported column: %T", v))
	}

	return &StringArrayColumn{
		Column:     make([][]string, 0, numRow),
		elem:       columnar,
		stringElem: stringElem,
		lcElem:     lcElem,
	}
}

func (c *StringArrayColumn) Reset(numRow int) {
	if cap(c.Column) >= numRow {
		c.Column = c.Column[:0]
	} else {
		c.Column = make([][]string, 0, numRow)
	}
}

func (c *StringArrayColumn) Type() reflect.Type {
	return stringSliceType
}

func (c *StringArrayColumn) Set(v any) {
	c.Column = v.([][]string)
}

func (c *StringArrayColumn) Value() any {
	return c.Column
}

func (c *StringArrayColumn) Nullable(nulls UInt8Column) any {
	panic("not implemented")
}

func (c *StringArrayColumn) Len() int {
	return len(c.Column)
}

func (c *StringArrayColumn) Index(idx int) any {
	return c.Column[idx]
}

func (c StringArrayColumn) Slice(s, e int) any {
	return c.Column[s:e]
}

func (c *StringArrayColumn) ConvertAssign(idx int, v reflect.Value) error {
	v.Set(reflect.ValueOf(c.Column[idx]))
	return nil
}

func (c *StringArrayColumn) AppendValue(v reflect.Value) {
	c.Column = append(c.Column, v.Interface().([]string))
}

func (c *StringArrayColumn) AppendPointer(typ reflect.Type, ptr unsafe.Pointer) {
	c.Column = append(c.Column, *(*[]string)(ptr))
}

func (c *StringArrayColumn) ReadFrom(rd *chproto.Reader, numRow int) error {
	if numRow == 0 {
		return nil
	}

	if cap(c.Column) >= numRow {
		c.Column = c.Column[:numRow]
	} else {
		c.Column = make([][]string, numRow)
	}

	if c.lcElem != nil {
		if err := c.lcElem.readPrefix(rd, numRow); err != nil {
			return err
		}
	}

	offsets := make([]int, numRow)

	for i := 0; i < len(offsets); i++ {
		offset, err := rd.UInt64()
		if err != nil {
			return err
		}
		offsets[i] = int(offset)
	}

	if err := c.elem.ReadFrom(rd, offsets[len(offsets)-1]); err != nil {
		return err
	}

	var prev int
	for i, offset := range offsets {
		c.Column[i] = c.stringElem.Column[prev:offset]
		prev = offset
	}

	return nil
}

func (c *StringArrayColumn) WriteTo(wr *chproto.Writer) error {
	if c.lcElem != nil {
		c.lcElem.writePrefix(wr)
	}

	_ = c.WriteOffset(wr, 0)
	return c.WriteData(wr)
}

var _ ArrayColumnar = (*StringArrayColumn)(nil)

func (c *StringArrayColumn) WriteOffset(wr *chproto.Writer, offset int) int {
	for _, el := range c.Column {
		offset += len(el)
		wr.UInt64(uint64(offset))
	}
	return offset
}

func (c *StringArrayColumn) WriteData(wr *chproto.Writer) error {
	for _, ss := range c.Column {
		c.stringElem.Column = ss
		if err := c.elem.WriteTo(wr); err != nil {
			return err
		}
	}
	return nil
}

//------------------------------------------------------------------------------

type ArrayLCStringColumn struct {
	*LCStringColumn
}

func (c ArrayLCStringColumn) Type() reflect.Type {
	return stringSliceType
}

func (c *ArrayLCStringColumn) WriteTo(wr *chproto.Writer) error {
	c.writeData(wr)
	return nil
}

func (c *ArrayLCStringColumn) ReadFrom(rd *chproto.Reader, numRow int) error {
	if numRow == 0 {
		return nil
	}
	return c.readData(rd, numRow)
}

//------------------------------------------------------------------------------

type GenericArrayColumn struct {
	Column reflect.Value

	typ       reflect.Type
	elem      Columnar
	arrayElem ArrayColumnar
}

var _ Columnar = (*GenericArrayColumn)(nil)

func NewGenericArrayColumn(typ reflect.Type, chType string, numRow int) Columnar {
	elemType := chArrayElemType(chType)
	if elemType == "" {
		panic(fmt.Errorf("invalid array type: %q (Go type is %s)",
			chType, typ.String()))
	}

	elem := NewColumn(typ.Elem(), elemType, 0)
	var arrayElem ArrayColumnar

	if _, ok := elem.(*LCStringColumn); ok {
		panic("not reached")
	}
	arrayElem, _ = elem.(ArrayColumnar)

	c := &GenericArrayColumn{
		typ:       reflect.SliceOf(typ),
		elem:      elem,
		arrayElem: arrayElem,
	}

	c.Column = reflect.MakeSlice(c.typ, 0, numRow)

	return c
}

func (c GenericArrayColumn) Type() reflect.Type {
	return c.typ.Elem()
}

func (c *GenericArrayColumn) Reset(numRow int) {
	if c.Column.Cap() >= numRow {
		c.Column = c.Column.Slice(0, 0)
	} else {
		c.Column = reflect.MakeSlice(c.typ, 0, numRow)
	}
}

func (c *GenericArrayColumn) Set(v any) {
	c.Column = reflect.ValueOf(v)
}

func (c *GenericArrayColumn) Value() any {
	return c.Column.Interface()
}

func (c *GenericArrayColumn) Nullable(nulls UInt8Column) any {
	panic("not implemented")
}

func (c *GenericArrayColumn) Len() int {
	return c.Column.Len()
}

func (c *GenericArrayColumn) Index(idx int) any {
	return c.Column.Index(idx).Interface()
}

func (c GenericArrayColumn) Slice(s, e int) any {
	return c.Column.Slice(s, e).Interface()
}

func (c *GenericArrayColumn) ConvertAssign(idx int, v reflect.Value) error {
	v.Set(c.Column.Index(idx))
	return nil
}

func (c *GenericArrayColumn) AppendValue(v reflect.Value) {
	c.Column = reflect.Append(c.Column, v)
}

func (c *GenericArrayColumn) AppendPointer(typ reflect.Type, ptr unsafe.Pointer) {
	c.AppendValue(reflect.NewAt(typ.Elem(), ptr).Elem())
}

func (c *GenericArrayColumn) ReadFrom(rd *chproto.Reader, numRow int) error {
	if c.Column.Cap() >= numRow {
		c.Column = c.Column.Slice(0, numRow)
	} else {
		c.Column = reflect.MakeSlice(c.typ, numRow, numRow)
	}

	if numRow == 0 {
		return nil
	}

	offsets := make([]int, numRow)
	for i := 0; i < len(offsets); i++ {
		offset, err := rd.UInt64()
		if err != nil {
			return err
		}
		offsets[i] = int(offset)
	}

	if err := c.elem.ReadFrom(rd, offsets[len(offsets)-1]); err != nil {
		return err
	}

	var prev int
	for i, offset := range offsets {
		c.Column.Index(i).Set(reflect.ValueOf(c.elem.Slice(prev, offset)))
		prev = offset
	}

	return nil
}

func (c *GenericArrayColumn) WriteTo(wr *chproto.Writer) error {
	_ = c.WriteOffset(wr, 0)

	colLen := c.Column.Len()
	for i := 0; i < colLen; i++ {
		// TODO: add SetValue or SetPointer
		c.elem.Set(c.Column.Index(i).Interface())

		var err error
		if c.arrayElem != nil {
			err = c.arrayElem.WriteData(wr)
		} else {
			err = c.elem.WriteTo(wr)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *GenericArrayColumn) WriteOffset(wr *chproto.Writer, offset int) int {
	colLen := c.Column.Len()

	for i := 0; i < colLen; i++ {
		el := c.Column.Index(i)
		offset += el.Len()
		wr.UInt64(uint64(offset))
	}

	if c.arrayElem == nil {
		return offset
	}

	offset = 0
	for i := 0; i < colLen; i++ {
		el := c.Column.Index(i)
		c.elem.Set(el.Interface()) // Use SetValue or SetPointer
		offset = c.arrayElem.WriteOffset(wr, offset)
	}

	return offset
}
