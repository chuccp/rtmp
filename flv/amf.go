package flv

import (
	"bytes"
	"encoding/binary"
	"github.com/chuccp/rtmp/util"
	"github.com/chuccp/utils/io"
	"github.com/chuccp/utils/log"
)

type AmfType byte

const (
	Number AmfType = iota
	Boolean
	String
	Object
	MovieClip
	Null
	Undefined
	Reference
	MixedArray = 8
	EndOfObject
	Array
	Date
	LongString
	Unsupported
	Recordset
	XML
	TypedObject
	AMF3data
)

func (t AmfType) Size() int {
	if t == Number {
		return 8
	}
	if t == Boolean {
		return 1
	}
	if t == String {
		return 2
	}
	if t == MixedArray {
		return 4
	}

	return 0
}

func (t AmfType) ReadValue(readStream *io.ReadStream) ([]byte, error) {

	if t == Number || t == Boolean {
		bytes, err := readStream.ReadBytes(t.Size())
		if err != nil {
			return nil, err
		}
		return bytes, err
	}
	if t == String {
		bytes, err := readStream.ReadBytes(t.Size())
		if err != nil {
			return nil, err
		}
		num := util.BigEndian(bytes)
		uintBytes, err := readStream.ReadUintBytes(num)
		if err != nil {
			return nil, err
		} else {
			return uintBytes, nil
		}
	}
	return nil, nil
}


type Amf struct {
	amf1Type   AmfType
	amf1Size   uint32
	amf1Name   string
	amf2Type   AmfType
	amf2Count  uint32
	parameters Parameters
}

func CreateAmf(name string, parameters Parameters) *Amf {
	return &Amf{amf1Name: name, parameters: parameters}
}

func (a *Amf) ToBytes() []byte {

	buff := new(bytes.Buffer)
	buff.WriteByte(2)

	nLen := len(a.amf1Name)
	buff.WriteByte(byte(nLen >> 8))
	buff.WriteByte(byte(nLen))
	buff.Write([]byte(a.amf1Name))

	buff.WriteByte(8)
	mixNum := len(a.parameters)
	buff.WriteByte(byte(mixNum >> 24))
	buff.WriteByte(byte(mixNum >> 16))
	buff.WriteByte(byte(mixNum >> 8))
	buff.WriteByte(byte(mixNum))

	for k, v := range a.parameters {
		kLen := len(k)
		buff.WriteByte(byte(kLen >> 8))
		buff.WriteByte(byte(kLen))
		buff.Write([]byte(k))
		if v.AmfType == Number {
			buff.WriteByte(0)
			buff.Write(v.Value)
		}else if v.AmfType==String{
			buff.WriteByte(2)
			nLen :=len(v.Value)
			buff.WriteByte(byte(nLen >> 8))
			buff.WriteByte(byte(nLen))
			buff.Write(v.Value)
		}
	}
	return nil
}

func NewAmf() *Amf {
	return &Amf{parameters: make(map[string]*Parameter)}
}

func (a *Amf) ReadAMF(readStream *io.ReadStream) error {
	b, err := readStream.ReadByte()
	if err != nil {
		return err
	}
	a.amf1Type = AmfType(b)
	bs, err := readStream.ReadBytes(a.amf1Type.Size())
	if err != nil {
		return err
	}
	a.amf1Size = util.BigEndian(bs)
	bytes, err := readStream.ReadUintBytes(a.amf1Size)
	if err != nil {
		return err
	}
	a.amf1Name = string(bytes)
	b, err = readStream.ReadByte()
	if err != nil {
		return err
	}
	a.amf2Type = AmfType(b)
	readBytes, err := readStream.ReadBytes(a.amf2Type.Size())
	if err != nil {
		return err
	}
	a.amf2Count = util.BigEndian(readBytes)
	return nil
}
func (a *Amf) ReadParam(key string) *Parameter {
	return a.parameters[key]
}
func (a *Amf) ReadParams(readStream *io.ReadStream) error {
	for i := 0; i < int(a.amf2Count); i++ {
		data, err := readStream.ReadBytes(2)
		if err != nil {
			return err
		}
		nameLen := binary.BigEndian.Uint16(data)
		bytes, err := readStream.ReadUintBytes(uint32(nameLen))
		if err != nil {
			return err
		}
		readByte, err := readStream.ReadByte()
		if err != nil {
			return err
		}
		type_ := AmfType(readByte)
		log.Info("key:", string(bytes))
		readBytes, err := type_.ReadValue(readStream)
		if err != nil {
			return err
		}
		np := NewParameter(string(bytes), type_, readBytes)
		a.parameters[np.Name] = np
	}
	return nil
}
