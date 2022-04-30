//go:build !amd64 && !arm64

package chschema

import (
	"github.com/uptrace/go-clickhouse/ch/chproto"
)

{{- range . }}

func (c *{{ .CHType }}Column) ReadFrom(rd *chproto.Reader, numRow int) error {
	c.Alloc(numRow)

	for i := range c.Column {
		n, err := rd.{{ .CHType }}()
		if err != nil {
			return err
		}
		c.Column[i] = n
	}

	return nil
}

func (c *{{ .CHType }}Column) WriteTo(wr *chproto.Writer) error {
	for _, n := range c.Column {
		wr.{{ .CHType }}(n)
	}
	return nil
}

{{- end }}
