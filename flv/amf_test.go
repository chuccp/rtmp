package flv

import (
	"fmt"
	"github.com/chuccp/rtmp/util"
	"testing"
)

func TestParameter_GetNumber(t *testing.T) {

	var data []byte = []byte{0x40,0x86,0x80,0,0,0,0,0}
	println(util.BytesToFloat64(data))
	println(util.BytesToFloat64(util.Float64ToBytes(util.BytesToFloat64(data))))

	println(fmt.Printf("%x!!!!",util.Float64ToBytes(util.BytesToFloat64(data))))

}
