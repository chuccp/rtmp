package util

func BigEndian(data []byte)uint32  {
	if len(data)==1{
		return uint32(data[0])
	}else if len(data)==2{
		return uint32(data[0])<<8|uint32(data[1])
	}else if len(data)==3{
		return uint32(data[0])<<16|uint32(data[1])<<8|uint32(data[2])
	}else {
		return uint32(data[0])<<24|uint32(data[1])<<16|uint32(data[2])<<8|uint32(data[3])
	}
}
