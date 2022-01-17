package micro

import (
	"github.com/hashicorp/go-msgpack/codec"
	"net"
)

func (c *Client) NewRPCMsgpackClient() (*msgpackClient, error) {
	if c.serviceErr !=nil{
		return nil, c.serviceErr
	}

	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return nil, err
	}
	return &msgpackClient{
		addr: c.addr,
		conn: conn,
		h:    &codec.MsgpackHandle{},
	}, nil
}

func (c *Client) NewRPCCodecClient() (*codecClient, error) {
	if c.serviceErr !=nil{
		return nil, c.serviceErr
	}
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return nil, err
	}
	return &codecClient{
		addr: c.addr,
		conn: conn,
		h:    &codec.BincHandle{},
	}, nil
}
