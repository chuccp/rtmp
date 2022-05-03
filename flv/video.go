package flv

import "github.com/chuccp/utils/io"

type Video struct {
	size uint32
	FrameType byte
	CodecID byte
	VideoData []byte
	readStream *io.ReadStream
}

func (v *Video) init()(*Video,error)  {

	readByte, err := v.readStream.ReadByte()
	if err != nil {
		return nil, err
	}else{
		v.FrameType = readByte>>4
		v.CodecID = readByte&0xF
	}
	v.VideoData,err = v.readStream.ReadUintBytes(v.size-1)
	return v,nil
}

func ParseVideo(tag *Tag) (*Video, error) {
	return (&Video{readStream: io.NewReadBytesStream(tag.Data),size: tag.DataSize}).init()
}