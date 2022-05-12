package h264

import (
	"bytes"
	videoError "github.com/chuccp/rtmp/error"
	"github.com/chuccp/rtmp/media"
	"github.com/chuccp/utils/io"
	"os"
)

type Decipher struct {
	reader   *io.ReadStream
	hasMatch bool
}

func NewDecipher() *Decipher {
	return &Decipher{}
}

func Open(path string) (*Decipher, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	} else {
		return &Decipher{reader: io.NewReadStream(file), hasMatch: false}, nil
	}
}
func (d *Decipher) Init(file *os.File)  {
	file.Seek(0,0)
	d.reader = io.NewReadStream(file)
	d.hasMatch = false
}

func (d *Decipher) Match() (bool, error) {
	data, err := d.reader.ReadBytes(3)
	if err != nil {
		return false, err
	} else {
		if bytes.Equal(data, []byte{0, 0, 1}) {
			return true, nil
		} else if bytes.Equal(data, []byte{0, 0, 0}) {
			b, err := d.reader.ReadByte()
			if err != nil {
				return false, err
			}
			if b == 1 {
				return true, nil
			}
		}
	}
	return false, videoError.UnknownFormatError
}
func (d *Decipher) DumpInfo() (*media.VideoInfo, error) {

	naul, err := d.ReadNAUL()
	if err != nil {
		return nil, err
	}
	if naul.IsSPS() {
		sps, err := ParseSPS(naul)
		if err != nil {
			return nil, err
		}
		vi:=media.NewVideoInfo(media.H264)
		width := (sps.picWidthInMbsMintus1+1)*16
		height:=(sps.picHeightInMapUnitsMinus1+1)*16
		vi.Width = uint32(width)
		vi.Height = uint32(height)
		vi.Framerate = sps.timeScale/sps.numUnitInTick/2
		return vi,nil
	}
	return nil, videoError.VideoFormatError
}

func (d *Decipher) ReadNAUL() (*NAUL, error) {
	if !d.hasMatch {
		d.hasMatch = true
		flag, err := d.Match()
		if err != nil {
			return nil, err
		}
		if !flag {
			return nil, videoError.UnknownFormatError
		}
	}
	buff := new(bytes.Buffer)
	var err0 error
	for {
		b0, err := d.reader.ReadByte()
		if err == nil {
			if b0 == 0 {
				b1, err := d.reader.ReadByte()
				if err == nil {
					if b1 == 0 {
						b2, err := d.reader.ReadByte()
						if err == nil {
							if b2 == 1 {
								break
							} else if b2 == 0 {
								b3, err := d.reader.ReadByte()
								if err == nil {
									if b3 == 1 {
										break
									} else {
										buff.WriteByte(b0)
										buff.WriteByte(b1)
										buff.WriteByte(b2)
										buff.WriteByte(b3)
									}
								} else {
									err0 = err
									break
								}
							} else {
								buff.WriteByte(b0)
								buff.WriteByte(b1)
								buff.WriteByte(b2)
							}
						} else {
							err0 = err
							break
						}
					} else {
						buff.WriteByte(b0)
						buff.WriteByte(b1)
					}
				} else {
					err0 = err
					break
				}
			} else {
				buff.WriteByte(b0)
			}
		} else {
			err0 = err
			break
		}
	}

	return NewNAUL(buff.Bytes()), err0
}
