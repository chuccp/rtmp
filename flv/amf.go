package flv

import (
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
func (t AmfType)ReadValue(readStream *io.ReadStream)([]byte,error){

	if t == Number || t==Boolean{
		bytes, err := readStream.ReadBytes(t.Size())
		if err != nil {
			return nil, err
		}
		return bytes, err
	}
	if t==String{
		bytes, err := readStream.ReadBytes(t.Size())
		if err != nil {
			return nil, err
		}
		num:=util.BigEndian(bytes)
		uintBytes, err := readStream.ReadUintBytes(num)
		if err != nil {
			return nil, err
		}else{
			return uintBytes, nil
		}
	}
	return nil, nil
}



type Parameters map[string]*Parameter

type Parameter struct {
	Name    string
	AmfType AmfType
	Value   []byte
}

func  NewParameter(Name string,AmfType AmfType,Value   []byte)*Parameter  {
	return &Parameter{Name:Name,AmfType: AmfType,Value: Value}
}

type Amf struct {
	amf1Type AmfType
	amf1Size uint32
	amf1Name string

	amf2Type  AmfType
	amf2Count uint32

	parameters Parameters

	readStream *io.ReadStream
}

func NewAmf(readStream *io.ReadStream) *Amf {
	return &Amf{readStream: readStream,parameters:make(map[string]*Parameter)}
}
func (a *Amf) ReadAMF() error {
	b, err := a.readStream.ReadByte()
	if err != nil {
		return err
	}
	a.amf1Type = AmfType(b)
	bs, err := a.readStream.ReadBytes(a.amf1Type.Size())
	if err != nil {
		return err
	}
	a.amf1Size = util.BigEndian(bs)
	bytes, err := a.readStream.ReadUintBytes(a.amf1Size)
	if err != nil {
		return err
	}
	a.amf1Name = string(bytes)
	b, err = a.readStream.ReadByte()
	if err != nil {
		return err
	}
	a.amf2Type = AmfType(b)
	readBytes, err := a.readStream.ReadBytes(a.amf2Type.Size())
	if err != nil {
		return err
	}
	a.amf2Count = util.BigEndian(readBytes)
	return nil
}
func (a *Amf) ReadParams() error{
	for i := 0; i < int(a.amf2Count); i++ {
		data,err:=a.readStream.ReadBytes(2)
		if err!=nil{
			return err
		}
		nameLen:=binary.BigEndian.Uint16(data)
		bytes, err := a.readStream.ReadUintBytes(uint32(nameLen))
		if err != nil {
			return err
		}
		readByte, err := a.readStream.ReadByte()
		if err != nil {
			return err
		}
		type_ :=AmfType(readByte)
		log.Info("key:",string(bytes))
		readBytes, err :=type_.ReadValue(a.readStream)
		if err != nil {
			return err
		}
		np:=NewParameter(string(bytes),type_,readBytes)
		a.parameters[np.Name] = np
	}
	return nil
}