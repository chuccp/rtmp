package flv

import (
	"encoding/binary"
	"github.com/chuccp/utils/io"
)

type Header struct {
	version byte
	flags byte
	headerSize uint32
}
func (h *Header) HasAudio()bool  {
	return h.flags&4==4
}
func (h *Header) HasVideo()bool  {
	return h.flags&1==1
}


func ParseHeader(reader *io.ReadStream) (*Header,error) {
	version,err := reader.ReadByte()
	if err!=nil{
		return nil, err
	}
	flags,err :=  reader.ReadByte()
	if err!=nil{
		return nil, err
	}
	data,err := reader.ReadBytes(4)
	if err!=nil{
		return nil, err
	}
	headerSize := binary.BigEndian.Uint32(data)
	return &Header{version: version,flags: flags,headerSize: headerSize},nil
}