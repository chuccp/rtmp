package h264

import (
	"bufio"
	"bytes"
	"os"
)

type Decipher struct {
	reader *bufio.Reader
	start bool
}

func Open(path string) (*Decipher,error) {
	file,err:=os.Open(path)
	if err!=nil{
		return nil, err
	}else{
		return &Decipher{reader: bufio.NewReader(file),start:false}, nil
	}
}

func (d *Decipher) ReadNAUL() (*NAUL,error) {
	buff:=new(bytes.Buffer)
	var err0 error
	for{
		b0,err:=d.reader.ReadByte()
		if err==nil{
			if b0==0{
				b1,err:=d.reader.ReadByte()
				if err==nil{
					if b1==0{
						b2,err:=d.reader.ReadByte()
						if err==nil{
							if b2==1{
								if !d.start{
									d.start = true
									continue
								}else{
									break
								}
							}else if b2==0{
								b3,err:=d.reader.ReadByte()
								if err==nil{
									if b3==1{
										if !d.start{
											d.start = true
											continue
										}else{
											break
										}
									}else{
										buff.WriteByte(b0)
										buff.WriteByte(b1)
										buff.WriteByte(b2)
										buff.WriteByte(b3)
									}
								}else{
									err0 = err
									break
								}
							}else{
								buff.WriteByte(b0)
								buff.WriteByte(b1)
								buff.WriteByte(b2)
							}
						}else{
							err0 = err
							break
						}
					}else{
						buff.WriteByte(b0)
						buff.WriteByte(b1)
					}
				}else{
					err0 = err
					break
				}
			}else{
				buff.WriteByte(b0)
			}
		}else{
			err0 = err
			break
		}
	}



	 return NewNAUL(buff.Bytes()), err0
}
