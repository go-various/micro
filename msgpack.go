package micro

import (
	"github.com/hashicorp/go-msgpack/codec"
	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
	"net"
	"sync"
	"time"
)

type msgpackClient struct {
	locker   sync.Mutex
	addr     string
	conn     net.Conn
	deadline time.Duration
	h        *codec.MsgpackHandle
}

func (c *msgpackClient) WithHandler(h *codec.MsgpackHandle) *msgpackClient {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.h = h
	return c
}

func (c *msgpackClient) WithDeadline(duration time.Duration) *msgpackClient {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.deadline = duration
	return c
}

func (c *msgpackClient) Call(method string, args interface{}, reply interface{}) error {
	if c.deadline > 0 {
		if err := c.conn.SetDeadline(time.Now().Add(c.deadline)); err != nil {
			return err
		}
	}
	defer c.conn.SetDeadline(time.Time{})

	rpcCodec := msgpackrpc.NewCodecFromHandle(true, true, c.conn, c.h)
	return msgpackrpc.CallWithCodec(rpcCodec, method, args, reply)
}

func (c *msgpackClient) Close() error {
	return c.conn.Close()
}
