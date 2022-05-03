package media

type IDecipher interface {
	Match()(bool,error)
	DumpInfo()*VideoInfo
}

