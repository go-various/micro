package micro

import "time"

// Client is 类似于工厂方法的实现
type Client struct {
	//经过lb逻辑后返回的服务地址
	addr string
	//传入的Service.GetServers接口的调用失败信息
	//执行逻辑前应该始终检查该错误
	serviceErr error

	hooks []Hook
}

func (c *Client) AddHooks(hooks ...Hook) {
	c.hooks = append(c.hooks, hooks...)
}

func (c *Client) Addr() string {
	return c.addr
}

func (c *Client) RestyClient() (*RestyClient, error) {
	if c.serviceErr != nil {
		return nil, c.serviceErr
	}
	cli := DefaultResty(c.addr)
	cli.hooks = c.hooks
	return cli, nil
}

func (c *Client) NewRestyClient(timeout time.Duration, InsecureSkipVerify bool) (*RestyClient, error) {
	if c.serviceErr != nil {
		return nil, c.serviceErr
	}
	cli := NewResty(c.addr, timeout, InsecureSkipVerify)
	cli.hooks = c.hooks
	return cli, nil
}
