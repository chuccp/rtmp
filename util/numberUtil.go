package util

import (
	"encoding/binary"
	"math"
)

func BigEndian(data []byte) uint32 {
	if len(data) == 1 {
		return uint32(data[0])
	} else if len(data) == 2 {
		return uint32(data[0])<<8 | uint32(data[1])
	} else if len(data) == 3 {
		return uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2])
	} else {
		return uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
	}
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64ToBytes(v float64) []byte {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, math.Float64bits(v))
	return data
}
