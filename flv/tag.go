package flv

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/chuccp/rtmp/h264"
	"github.com/chuccp/utils/io"
)

type Tag struct {
	Type      byte
	DataSize  uint32
	Timestamp uint32
	StreamID  uint32
	Data      []byte
	TagSize   uint32
	reader    *bufio.Reader
}

func (t *Tag) IsAudio() bool {
	return t.Type&0x08 == 0x08
}
func (t *Tag) IsVideo() bool {
	return t.Type&0x09 == 0x09
}
func (t *Tag) IsScript() bool {
	return t.Type&0x12 == 0x12
}



func NewTag() *Tag {
	return &Tag{}
}
func ReadTag(reader *io.ReadStream) (*Tag, error) {
	tag := &Tag{}
	var err error
	tag.Type, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}
	b, err := reader.ReadBytes(3)
	if err != nil {
		return nil, err
	}

	tag.DataSize = uint32(b[2]) | uint32(b[1])<<8 | uint32(b[0])<<16
	t, err := reader.ReadBytes(3)
	if err != nil {
		return nil, err
	}
	tag.Timestamp = uint32(t[2]) | uint32(t[1])<<8 | uint32(t[0])<<16
	timestampEx, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	tag.Timestamp = tag.Timestamp | uint32(timestampEx)<<24
	s, err := reader.ReadBytes(3)
	if err != nil {
		return nil, err
	}
	tag.StreamID = uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])

	tag.Data, err = reader.ReadUintBytes(tag.DataSize)
	if err != nil {
		return nil, err
	}

	dd, err := reader.ReadBytes(4)
	if err != nil {
		return nil, err
	}
	tag.TagSize = binary.BigEndian.Uint32(dd)
	return tag, nil
}

func CreateVideoInfoTag(sps *h264.SPS,pps []byte,time uint32) []byte {

	var buff = new(bytes.Buffer)
	buff.WriteByte(0x17)
	buff.Write([]byte{0,0,0,0})

	buff.WriteByte(0x1)
	sps.ProfileIdc()
	buff.WriteByte(byte(sps.ProfileIdc()))
	buff.WriteByte(0)
	buff.WriteByte(byte(sps.LevelIdc()))
	buff.WriteByte(0xff)
	buff.WriteByte(0xE1)

	data:=sps.Bytes()
	l:=len(data)
	buff.WriteByte(byte(l >> 8))
	buff.WriteByte(byte(l))
	buff.Write(data)
	buff.WriteByte(1)

	pl:=len(pps)
	buff.WriteByte(byte(pl >> 8))
	buff.WriteByte(byte(pl))
	buff.Write(pps)
	return buff.Bytes()
}
func CreateVideoInfoTag2(SEI []byte,IDR []byte,time uint32)[]byte {
	var buff = new(bytes.Buffer)
	buff.WriteByte(0x17)
	buff.WriteByte(1)

	return nil
}