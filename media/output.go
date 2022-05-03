package media

import "os"

type OutputVideo struct {
	file *os.File
	filePath string
	avCodeId AVCodeId
}

func (v *OutputVideo) Write(vp *VideoPacket)  {


}

func CreateOutputVideo(filePath string)(*OutputVideo,error)  {
	return &OutputVideo{filePath:filePath},nil
}