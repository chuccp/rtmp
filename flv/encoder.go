package flv

import (
	"github.com/chuccp/utils/io"
	"os"
)

type Encoder struct {
	write *io.WriteStream
}

func (e *Encoder) WriteHeaderAndZeroTag(audio bool, video bool) {
	var b byte = 0x0
	if audio {
		b = b | 4
	}
	if video {
		b = b | 1
	}
	e.write.Write([]byte{0x46, 0x4C, 0x56, 0x01, b, 0, 0, 0, 9, 0, 0, 0, 0})
}
func (e *Encoder) WriteScript() {

	var parameters = make(Parameters)
	parameters.Add(NumberParameter(WIDTH, 720))
	parameters.Add(NumberParameter(DURATION, 14.415))
	parameters.Add(NumberParameter(HEIGHT, 1280))
	CreateAmf("onMetaData", parameters)

}

func Create(path string) (*Encoder, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	} else {
		return &Encoder{write: io.NewWriteStream(file)}, nil
	}
}
