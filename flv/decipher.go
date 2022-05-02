package flv

import (
	"bufio"
	"encoding/binary"
	"github.com/chuccp/utils/io"
	"os"
)

type Decipher struct {
	reader	*io.ReadStream
	start bool
}

func (d *Decipher) ReadHeader() (*Header,error) {
	return PasreHeader(d.reader)
}

func (d *Decipher) ReadZeroTag() (uint32,error) {
	data,err:=d.reader.ReadBytes(4)
	if err!=nil{
		return 0, err
	}
	return binary.BigEndian.Uint32(data),nil
}
func (d *Decipher) ReadTag() (*Tag,error) {
	return ReadTag(d.reader)
}



func Open(path string) (*Decipher,error) {
	file,err:=os.Open(path)
	if err!=nil{
		return nil, err
	}else{
		return &Decipher{reader:io.NewReadStream( bufio.NewReader(file)),start:false}, nil
	}
}
