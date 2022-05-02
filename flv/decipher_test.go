package flv

import (
	"fmt"
	"testing"
)

func TestDecipher(t *testing.T) {

	d, err := Open("C:\\Users\\cooge\\Videos\\123321.flv")
	if err != nil {
		return
	}else{
		header,err:=d.ReadHeader()
		if err==nil{
			t.Log(fmt.Printf("signature %x",header.signature))
			t.Log(fmt.Printf("version %x",header.version))
			t.Log(header.HasAudio())
			t.Log(header.HasVideo())
			t.Log(header.headerSize)
			v,err:=d.ReadZeroTag()
			if err==nil{
				t.Log("v",v)
				var num = 0
				for{
					tag,err :=	d.ReadTag()
					if err==nil{
						if tag.IsScript(){
							t.Log("===","Script")

							script, err := ParseScript(tag)
							if err != nil {
								return
							}else{
								amf:=script.Amf()
								t.Log(amf.amf1Name)
							}
						}else
						if tag.IsVideo(){
							t.Log("===","video")
						}else
						if tag.IsAudio(){
							t.Log("===","audio")
						}else{
							t.Log("===","error")
						}
					}else{
						break
					}
					num++
					if num>5{
						break
					}
				}



			}
		}
	}
}
