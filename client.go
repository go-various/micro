package micro

import (
	"github.com/hashicorp/go-msgpack/codec"
	"net"
	"time"
)

// Client is 类似于工厂方法的实现
type Client struct {
	//经过lb逻辑后返回的服务地址
	addr       string

	//传入的Service.GetServers接口的调用失败信息
	//执行逻辑前应该始终检查该错误
	serviceErr error
}

func (c *Client)Addr()string  {
	return c.addr
}

func (c *Client) RestyClient() (*RestyClient,error) {
	if c.serviceErr != nil{
		return nil, c.serviceErr
	}
	cli := DefaultResty(c.addr)
	return cli, nil
}

func (c *Client) NewRestyClient(timeout time.Duration, InsecureSkipVerify bool) (*RestyClient, error) {
	if c.serviceErr !=nil{
		return nil, c.serviceErr
	}

	return NewResty(c.addr, timeout, InsecureSkipVerify), nil
}

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
