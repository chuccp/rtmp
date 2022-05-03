package main

import "github.com/chuccp/rtmp/media"

func main() {

	inputVideo, err := media.OpenVideo("C:\\Users\\cooge\\Videos\\123321.h264")
	if err != nil {
		return
	}
	videoInfo := inputVideo.ReadVideoInfo()
	if videoInfo!=nil{


	}


}
