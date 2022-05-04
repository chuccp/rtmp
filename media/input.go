package media

import "os"

type InputVideo struct {
	file *os.File
	avCodeId AVCodeId
	decipher IDecipher

}
func (v *InputVideo) ReadVideoInfo() (*VideoInfo,error) {

	var err error
	v.decipher, err = GetDecipher(v.file)
	if err != nil {
		return nil, err
	}
	v.file.Seek(0,0)
	v.decipher.Init(v.file)
	return v.decipher.DumpInfo()
}

func OpenVideo(file string) (*InputVideo,error) {
	fl, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return &InputVideo{file:fl},nil
}
