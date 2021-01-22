package micro

import (
	"github.com/hashicorp/go-msgpack/codec"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type codecClient struct {
	locker sync.Mutex
	addr string
	conn net.Conn
	h codec.Handle
	deadline time.Duration
}


func (c *codecClient)WithHandler(h codec.Handle) *codecClient {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.h = h
	return c
}

func (c *codecClient) WithDeadline(duration time.Duration)*codecClient {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.deadline = duration
	return c
}

func (c *codecClient) Call(method string, args interface{}, reply interface{})error  {
	if c.deadline > 0{
		if err := c.conn.SetDeadline(time.Now().Add(c.deadline)); err != nil {
			return err
		}
	}
	defer c.conn.SetDeadline(time.Time{})

	rpcCodec := codec.GoRpc.ClientCodec(c.conn, c.h)
	client := rpc.NewClientWithCodec(rpcCodec)
	return client.Call(method, args, reply)
}



func (c *codecClient) Close()error  {
	return c.conn.Close()
}