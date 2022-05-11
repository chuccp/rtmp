package flv

import "github.com/chuccp/rtmp/util"

const (
	WIDTH         = "width"
	HEIGHT        = "height"
	DURATION      = "duration"
	VIDEODATARATE = "videodatarate"
	FRAMERATE     = "framerate"
	VIDEOCODECID  = "videocodecid"
)

type Parameters map[string]*Parameter

func (ps Parameters) Add(p *Parameter) {
	ps[p.Name] = p
}

type Parameter struct {
	Name    string
	AmfType AmfType
	Value   []byte
}

func StringParameter(name string, value string) *Parameter {
	p := &Parameter{}
	p.AmfType = String
	p.Name = name
	p.Value = []byte(value)
	return p

}
func NumberParameter(name string, value float64) *Parameter {
	p := &Parameter{}
	p.AmfType = Number
	p.Name = name
	p.Value = util.Float64ToBytes(value)
	return p

}

func NewParameter(Name string, AmfType AmfType, Value []byte) *Parameter {
	return &Parameter{Name: Name, AmfType: AmfType, Value: Value}
}
