package h264

import "fmt"

type NAULType byte

const (
	SPSType NAULType = 7
	PPSType          = 8
	SEIType          = 6
)

type NAUL struct {
	data          []byte
	header byte
	forbiddenZero bool
	nalRefIdc     byte
	nalUnitType   NAULType
}

func (n *NAUL) init() *NAUL {

	n.forbiddenZero = n.header>>7 == 0
	n.nalRefIdc = n.header & 127 >> 5
	n.nalUnitType = NAULType(n.header & 31)
	return n

}
func (n *NAUL) ToHexString() string {
	return fmt.Sprintf(" %x ", n.data)
}
func (n *NAUL) NAULType() NAULType {
	return n.nalUnitType
}

func NewNAUL(data []byte) *NAUL {
	return (&NAUL{header:data[0],data: data[1:]}).init()
}
