package main

import "github.com/chuccp/rtmp/flv"

func main() {

	//dv,err:=flv.Open("C:\\Users\\cooge\\Videos\\123321.flv")
	//if err==nil{
	//
	//
	//
	//}

	v:=flv.AmfType(0x2)
	println(v.Size())
}
