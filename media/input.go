package media

import "os"

type InputVideo struct {
	file *os.File
	avCodeId AVCodeId

}
func (v *InputVideo) ReadVideoInfo() *VideoInfo {





	return &VideoInfo{}
}

func OpenVideo(file string) (*InputVideo,error) {
	fl, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return &InputVideo{file:fl},nil
}
