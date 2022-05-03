package flv

import (
	"bufio"
	"encoding/binary"
	videoError "github.com/chuccp/rtmp/error"
	"github.com/chuccp/utils/io"
	"os"
)

type Decipher struct {
	reader	*io.ReadStream
	start bool
	signature []byte
}
func (d *Decipher)signatureMatch()(bool,error){
	var err error
	d.signature,err = d.reader.ReadBytes(3)
	if err!=nil{
		return false, err
	}else{
		if string(d.signature) == "FLV"{
			return true, nil
		}
		return false, videoError.VideoFormatError
	}
}

func (d *Decipher) ReadHeader() (*Header,error) {
	if d.signature==nil{
		_,err:=d.signatureMatch()
		if err!=nil{
			return nil, err
		}
	}
	return ParseHeader(d.reader)
}

func (d *Decipher) ReadZeroTag() (uint32,error) {
	data,err:= d.reader.ReadBytes(4)
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
