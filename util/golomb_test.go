package util

import (
	"bytes"
	"testing"
)

func TestGolomb(t *testing.T) {


	var buff = new(bytes.Buffer)
	buff.WriteByte(217)
	buff.WriteByte(0x0F)
	buff.WriteByte(0x0F)
	buff.WriteByte(0x0F)
	buff.WriteByte(0x0F)
	gd:=NewGolombDecode(buff.Bytes())

	v,err:=gd.ReadUGolomb()
	if err==nil{
		t.Logf("%x",v)
	}
	v,err=gd.ReadUGolomb()
	if err==nil{
		t.Logf("%x",v)
	}
	v,err=gd.ReadUGolomb()
	if err==nil{
		t.Logf("%x",v)
	}

}
