package micro

type Factory func(addrs Addrs)LoadBalance

type LoadBalance interface {
	Client() *Client
}

type Meta struct {
	Name string
	Value string
}

type Addrs map[string][]Meta

type lbClient struct {
	lb LoadBalance
}
func NewLBClient(addrs Addrs) *lbClient {
	return &lbClient{
		lb:    lb(addrs),
	}
}

func (lb *lbClient)LBClient()*Client {
	return lb.lb.Client()
}

