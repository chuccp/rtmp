package util

import (
	"bytes"
)

type GolombDecode struct {
	buff    *bytes.Buffer
	off     int
	cutByte byte
	start   bool
	sIndex  uint8
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

	err := g.reset()
	if err != nil {
		return false, err
	}

	v := g.cutByte<<g.off>>7 == 1
	g.off = g.off + 1
	return v, nil
}
func (g *GolombDecode) Read3Uint() (uint8, error) {

	var i uint8 = 0
	b, err := g.ReadBit()
	if err != nil {
		return 0, err
	}
	if b {
		i = i | 4
	}
	b, err = g.ReadBit()
	if err != nil {
		return 0, err
	}
	if b {
		i = i | 2
	}
	b, err = g.ReadBit()
	if err != nil {
		return 0, err
	}
	if b {
		i = i | 1
	}
	return i, nil
}
func (g *GolombDecode) ReadByte() (byte, error) {

	b, err := g.buff.ReadByte()
	if (g.sIndex == 0 || g.sIndex == 1) && b == 0 {
		g.sIndex++
	} else if g.sIndex == 2 && b == 3 {
		b, err = g.buff.ReadByte()
		g.sIndex = 0
		return b, err
	} else {
		g.sIndex = 0
	}
	return b, err
}
func (g *GolombDecode) reset() error {
	if g.off == 8 {
		g.off = 0
		b, err := g.ReadByte()
		if err != nil {
			return err
		} else {
			g.cutByte = b
		}
	}
	return nil
}

func (g *GolombDecode) ReadUint8() (uint8, error) {
	err := g.reset()
	if err != nil {
		return 0, err
	}
	u := g.cutByte << g.off
	b, err := g.ReadByte()
	if err != nil {
		return 0, err
	} else {
		g.cutByte = b
	}

	sLen := 8 - g.off
	return (b >> sLen) | u, nil
}
func (g *GolombDecode) ReadUint16() (uint16, error) {
	err := g.reset()
	if err != nil {
		return 0, err
	}

	u16 := uint16(g.cutByte) << (g.off + 8)
	b, err := g.ReadByte()
	if err != nil {
		return 0, err
	}
	sLen := 8 - g.off
	u16 = u16 + (uint16(b) << sLen)
	b, err = g.ReadByte()
	if err != nil {
		return 0, err
	} else {
		g.cutByte = b
	}
	return uint16(b>>sLen) | u16, nil
}
func (g *GolombDecode) ReadUint32() (uint32, error) {
	err := g.reset()
	if err != nil {
		return 0, err
	}

	u32 := uint32(g.cutByte) << (g.off + 24)
	b0, err := g.ReadByte()
	if err != nil {
		return 0, err
	}
	b1, err := g.ReadByte()
	if err != nil {
		return 0, err
	}
	b2, err := g.ReadByte()
	if err != nil {
		return 0, err
	}

	u32 = u32 | (uint32(b0) << (g.off + 16))
	u32 = u32 | (uint32(b1) << (g.off + 8))
	u32 = u32 | (uint32(b2) << (g.off))

	sLen := 8 - g.off
	b3, err := g.ReadByte()
	if err != nil {
		return 0, err
	} else {
		g.cutByte = b3
	}
	return uint32(b3>>sLen) | u32, nil
}

func (g *GolombDecode) readValue(len int) (uint64, error) {

	err := g.reset()
	if err != nil {
		return 0, err
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
				b, err := g.ReadByte()
				if err == nil {
					len = len - 8
					value = value + uint64(b)<<len
				} else {
					return 0, err
				}
			} else {
				b, err := g.ReadByte()
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
		v, err := g.ReadByte()
		if err != nil {
			return 0, err
		} else {
			g.cutByte = v
		}
	}
	if g.off == 8 {
		g.off = 0
		v, err := g.ReadByte()
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
		v, err := g.ReadByte()
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
