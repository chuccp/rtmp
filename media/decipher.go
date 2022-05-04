package media

import "os"

type IDecipher interface {
	Init(file *os.File)
	Match()(bool,error)
	DumpInfo()(*VideoInfo,error)
}

