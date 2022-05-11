package flv

import (
	"bufio"
	"encoding/binary"
	videoError "github.com/chuccp/rtmp/error"
	"github.com/chuccp/rtmp/media"
	"github.com/chuccp/rtmp/util"
	"github.com/chuccp/utils/io"
	"os"
)

type Decipher struct {
	reader	*io.ReadStream
	start bool
	hasMatch bool
	signature []byte
}
func (d *Decipher) Match()(bool,error){
	var err error
	signature,err := d.reader.ReadBytes(3)
	if err!=nil{
		return false, err
	}else{
		if string(signature) == "FLV"{
			return true, nil
		}
		return false, videoError.VideoFormatError
	}
}

func (d *Decipher) ReadHeader() (*Header,error) {
	if d.signature==nil{
		_,err:=d.Match()
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

func (d *Decipher) Init(file *os.File){
	d.reader = io.NewReadStream(file)
	d.hasMatch = false
}

func (d *Decipher) DumpInfo()(*media.VideoInfo,error){
	_,err:=d.ReadHeader()
	if err!=nil{
		return nil, err
	}
	_,err=d.ReadZeroTag()
	if err!=nil{
		return nil, err
	}
	tag,err:=d.ReadTag()
	if err!=nil{
		return nil, err
	}

	if tag.IsScript(){
		vi:=media.NewVideoInfo(media.H264)
		script, err := ParseScript(tag)
		if err != nil {
			return nil, err
		}
		amf:=script.Amf()
		width:=amf.ReadParam("width")
		height:=amf.ReadParam("height")
		duration:=amf.ReadParam("duration")
		framerate:=amf.ReadParam("framerate")
		vi.Width = uint32(util.BytesToFloat64(width.Value))
		vi.Height = uint32(util.BytesToFloat64(height.Value))
		vi.Duration = util.BytesToFloat64(duration.Value)
		vi.Framerate = uint32(util.BytesToFloat64(framerate.Value))
		return vi, nil
		
	}
	return nil, videoError.VideoFormatError
}


func NewDecipher() *Decipher {

	return &Decipher{}
}

func Open(path string) (*Decipher,error) {
	file,err:=os.Open(path)
	if err!=nil{
		return nil, err
	}else{
		return &Decipher{reader:io.NewReadStream( bufio.NewReader(file)),start:false}, nil
	}
}
