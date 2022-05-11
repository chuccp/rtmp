package flv

import (
	"bufio"
	"encoding/binary"
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
