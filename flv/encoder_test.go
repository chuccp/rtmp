package flv

import "testing"

func TestEncoder(t *testing.T) {

	encoder,err:=Create("C:\\Users\\cao\\Videos\\8988.flv")
	if err!=nil{
		t.Error(err)
	}else{
		encoder.WriteHeaderAndZeroTag(false,true)
	}
}