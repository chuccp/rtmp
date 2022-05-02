package flv

import (
	"github.com/chuccp/utils/io"
)

type Script struct {
	amf *Amf
}

func (s *Script) Amf() *Amf {
	return s.amf
}
func (s *Script) init() (*Script, error) {
	err := s.amf.ReadAMF()
	if err != nil {
		return nil, err
	}
	err = s.amf.ReadParams()
	return s, err
}
func ParseScript(tag *Tag) (*Script, error) {
	return (&Script{amf: NewAmf(io.NewReadBytesStream(tag.Data))}).init()
}
