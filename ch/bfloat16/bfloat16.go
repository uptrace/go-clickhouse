package bfloat16

import (
	"math"
)

type Map map[T]uint64

type T uint16

func From(f float64) T {
	return FromFloat32(float32(f))
}

func FromFloat32(f float32) T {
	return T(math.Float32bits(f) >> 16)
}

func (f T) Float32() float32 {
	return math.Float32frombits(uint32(f) << 16)
}

func (f T) Float64() float64 {
	return float64(f.Float32())
}
