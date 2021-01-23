package micro

import (
	"github.com/hashicorp/go-msgpack/codec"
	"net"
	"time"
)

// Client
type Client struct {
	addr string
}

func (c *Client) RestyClient() *RestyClient {
	return DefaultResty(c.addr)
}

func (c *Client) NewRestyClient(timeout time.Duration, InsecureSkipVerify bool) *RestyClient {
	return NewResty(c.addr, timeout, InsecureSkipVerify)
}

func (c *Client) NewRPCMsgpackClient() (*msgpackClient, error) {
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
