package media

import (
	"container/list"
	error2 "github.com/chuccp/rtmp/error"
	"os"
)

var decipherList = list.New()


func Add(decipher IDecipher)  {
	decipherList.PushFront(decipher)
}

func GetDecipher(file *os.File)(IDecipher,error) {
	for ele := decipherList.Front(); ele!=nil ; ele = ele.Next() {
		id:=(ele.Value).(IDecipher)
		flag, err := id.Match()
		if err!=nil{
			return nil, err
		}
		if flag{
			return id, nil
		}
	}
	return nil, error2.UnknownFormatError
}
