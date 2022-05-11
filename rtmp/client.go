package rtmp

import "github.com/chuccp/utils/io"

type Client struct {
	Host string
	Port int
	netStream *io.NetStream

}

func NewClient(Host string,Port int) *Client {
	return &Client{Host: Host,Port: Port}
}

func (c *Client) Start()(err error)  {
	xc:=io.NewXConn(c.Host,c.Port)
	c.netStream,err = xc.Create()
	return err
}
func (c *Client) handshake(){



}



