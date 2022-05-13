package flv

import (
	"bytes"
	"github.com/chuccp/rtmp/h264"
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
func (e *Encoder) WriteScript(width float64, height float64, duration float64) {

	var buff = new(bytes.Buffer)
	buff.WriteByte(18)

	var parameters = make(Parameters)
	parameters.Add(NumberParameter(WIDTH, width))
	parameters.Add(NumberParameter(DURATION, duration))
	parameters.Add(NumberParameter(HEIGHT, height))
	amf := CreateAmf("onMetaData", parameters)
	data := amf.ToBytes()
	dataSize := len(data)
	buff.WriteByte(byte(dataSize >> 16))
	buff.WriteByte(byte(dataSize >> 8))
	buff.WriteByte(byte(dataSize))
	buff.Write([]byte{0, 0, 0, 0, 0, 0, 0})

	buff.Write(data)
	buff.Write([]byte{0, 0, 9})

	wLen := uint32(buff.Len())
	buff.WriteByte(byte(wLen >> 24))
	buff.WriteByte(byte(wLen >> 16))
	buff.WriteByte(byte(wLen >> 8))
	buff.WriteByte(byte(wLen))

	e.write.Write(buff.Bytes())

}
func (e *Encoder) WriteVideoInfo(sps *h264.SPS, pps []byte) {
	var buff = new(bytes.Buffer)
	buff.WriteByte(9)
	data := CreateVideoInfoTag(sps, pps, 0)
	l := len(data)
	buff.WriteByte(byte(l >> 16))
	buff.WriteByte(byte(l >> 8))
	buff.WriteByte(byte(l))
	buff.Write([]byte{0, 0, 0, 0, 0, 0, 0})
	buff.Write(data)
	ln := buff.Len()
	buff.WriteByte(byte(ln >> 24))
	buff.WriteByte(byte(ln >> 16))
	buff.WriteByte(byte(ln >> 8))
	buff.WriteByte(byte(ln))
	e.write.Write(buff.Bytes())
}
func (e *Encoder) WriteSEIAndIDR(SEI []byte, IDR []byte){
	var buff = new(bytes.Buffer)
	buff.WriteByte(9)

}




func Create(path string) (*Encoder, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	} else {
		return &Encoder{write: io.NewWriteStream(file)}, nil
	}
}
