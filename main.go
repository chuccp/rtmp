package main

import (
	"github.com/chuccp/rtmp/h264"
	"github.com/chuccp/rtmp/media"
	"github.com/chuccp/utils/log"
)

func main() {
	media.Add(h264.NewDecipher())
	inputVideo, err := media.OpenVideo("C:\\Users\\cooge\\Videos\\123321.h264")
	if err != nil {
		return
	}

	info, err := inputVideo.ReadVideoInfo()
	if err != nil {
		return
	}
	log.Info(info)


}
