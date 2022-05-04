package main

import (
	"github.com/chuccp/rtmp/h264"
	"github.com/chuccp/rtmp/media"
	"log"
)

func main() {
	media.Add(h264.NewDecipher())
	inputVideo, err := media.OpenVideo("C:\\Users\\cooge\\Videos\\123321.h264")
	if err != nil {
		return
	}

	info, err := inputVideo.ReadVideoInfo()
	if err != nil {
		log.Panicln(err)
	}
	log.Println(info.Width,"=====",info.Height)


}
