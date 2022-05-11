package main

import (
	"github.com/chuccp/rtmp/flv"
	"github.com/chuccp/rtmp/h264"
	"github.com/chuccp/rtmp/media"
	"log"
)

func main() {
	media.Add(h264.NewDecipher())
	media.Add(flv.NewDecipher())
	inputVideo, err := media.OpenVideo("C:\\Users\\cao\\Videos\\123321.flv")
	if err != nil {
		log.Panicln(err)
		return
	}
	info, err := inputVideo.ReadVideoInfo()
	if err != nil {
		log.Panicln(err)
	}

	log.Println(info.Width,"=====",info.Height,"==",info.Duration,"==",info.Framerate)


}
