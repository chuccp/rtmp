package flv

import (
	"fmt"
	"github.com/chuccp/utils/log"
	"testing"
)
func TestDecipher(t *testing.T) {

	d, err := Open("C:\\Users\\cooge\\Videos\\123321.flv")
	if err != nil {
		return
	}else{
		header, err :=d.ReadHeader()
		if err==nil{
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
								t.Log(err)
								return
							}else{
								amf:=script.Amf()
								t.Log(amf.amf1Name)
							}
						}else
						if tag.IsVideo(){
							t.Log("===","video")

							v,err:=ParseVideo(tag)
							if err!=nil{
								t.Log(err)
								return
							}
							log.Info(v.FrameType," ",v.CodecID)

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
