package util

import (
	"bytes"
)

type GolombDecode struct {
	buff    *bytes.Buffer
	off     int
	cutByte byte
	start   bool
}

func (g *GolombDecode) ReadUGolomb() (uint64, error) {
	len, err := g.readValueLen()
	if err != nil {
		return 0, err
	} else {
		return g.readValue(len)
	}
}
func (g *GolombDecode) ReadGolomb() (int64, error) {
	len, err := g.readValueLen()
	if err != nil {
		return 0, err
	} else {
		v, err := g.readValue(len)
		if err != nil {
			return 0, err
		} else {
			v1 := v >> 1
			v2 := v & 1
			if v2 == 1 {
				return -int64(v1), nil
			}
			return int64(v1), nil
		}
	}
}
func (g *GolombDecode) ReadBit() (bool, error) {
	if g.off == 8 {
		g.off = 0
		b, err := g.buff.ReadByte()
		if err != nil {
			return false, err
		} else {
			g.cutByte = b
		}
	}
	v:=g.cutByte<<g.off>>7 == 1
	g.off = g.off + 1
	return v, nil

}

func (g *GolombDecode) readValue(len int) (uint64, error) {

	if g.off == 8 {
		g.off = 0
		b, err := g.buff.ReadByte()
		if err != nil {
			return 0, err
		}
		g.cutByte = b
	}

	sLen := 8 - g.off
	if sLen >= len {
		bb := g.cutByte << g.off >> (8 - len)
		g.off = g.off + len
		return uint64(bb) - 1, nil
	} else {
		var value uint64 = 0
		var v = g.cutByte << g.off >> g.off
		len = len - 8 + g.off
		value = value + uint64(v)<<len
		for {
			if len > 8 {
				b, err := g.buff.ReadByte()
				if err == nil {
					len = len - 8
					value = value + uint64(b)<<len
				} else {
					return 0, err
				}
			} else {
				b, err := g.buff.ReadByte()
				if err == nil {
					value = value + uint64(b)>>(8-len)
					g.off = len
					g.cutByte = b
					break
				} else {
					return 0, err
				}
			}
		}
		return value - 1, nil
	}

}
func (g *GolombDecode) readValueLen() (int, error) {
	len := 0
	if !g.start {
		g.start = true
		v, err := g.buff.ReadByte()
		if err != nil {
			return 0, err
		} else {
			g.cutByte = v
		}
	}
	if g.off == 8 {
		g.off = 0
		v, err := g.buff.ReadByte()
		if err != nil {
			return 0, err
		} else {
			g.cutByte = v
		}
	}
	for {
		num := readByteLen(g.cutByte, g.off)
		len = len + num
		if num+g.off != 8 {
			g.off = num + g.off
			break
		}
		v, err := g.buff.ReadByte()
		if err != nil {
			return 0, err
		} else {
			g.cutByte = v
			g.off = 0
		}
	}
	return len + 1, nil
}

func readByteLen(b byte, off int) int {
	b = b << off >> off
	for i := 7 - off; i >= 0; i-- {
		if (b >> i) != 0 {
			return 8 - i - 1 - off
		}
	}
	return 8 - off
}

func NewGolombDecode(data []byte) *GolombDecode {

	return &GolombDecode{buff: bytes.NewBuffer(data), off: 0, start: false}
}
