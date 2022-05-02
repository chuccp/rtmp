package flv

import (
	"encoding/binary"
	"github.com/chuccp/utils/io"
)

type Header struct {
	signature []byte
	version byte
	flags byte
	headerSize uint32
	reader *io.ReadStream
}

func (h *Header) HasAudio()bool  {
	return h.flags&4==4
}
func (h *Header) HasVideo()bool  {
	return h.flags&1==1
}

func (h *Header) init()(*Header,error) {
	var err error
	h.signature,err = h.reader.ReadBytes(3)
	if err!=nil{
		return nil, err
	}
	h.version,err = h.reader.ReadByte()
	if err!=nil{
		return nil, err
	}
	h.flags,err =  h.reader.ReadByte()
	if err!=nil{
		return nil, err
	}
	data,err := h.reader.ReadBytes(4)
	if err!=nil{
		return nil, err
	}
	h.headerSize = binary.BigEndian.Uint32(data)
	return h,nil
}
func PasreHeader(reader *io.ReadStream)(*Header,error)  {
	return (&Header{reader: reader}).init()
}