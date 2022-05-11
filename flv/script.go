package flv

import (
	"github.com/chuccp/utils/io"
)

type Script struct {
	amf *Amf
	readStream *io.ReadStream
}

func (s *Script) Amf() *Amf {
	return s.amf
}
func (s *Script) init() (*Script, error) {
	err := s.amf.ReadAMF(s.readStream)
	if err != nil {
		return nil, err
	}
	err = s.amf.ReadParams(s.readStream)
	return s, err
}
func ParseScript(tag *Tag) (*Script, error) {
	readStream := io.NewReadBytesStream(tag.Data)

	return (&Script{amf: NewAmf(),readStream:readStream}).init()
}
