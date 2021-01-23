package micro

type LBType string

const (
	LBTypeRandom           LBType = "lb:rnd"
	LBTypeRoundRobin       LBType = "lb:rr"
	LBTypeConsistentHash   LBType = "lb:hash"
	LBTypeLeastConnections LBType = "lb:lcs"
)

type Factory func(addrs Addrs) LoadBalance

type LoadBalance interface {
	Client() *Client
}

type Server struct {
	ID          string
	Address     string
	Weight      int
	TPSDelay    float64
	Connections int
}

type Addrs map[string][]Server

type lbClient struct {
	lb LoadBalance
}

func NewLBClient(addrs Addrs) *lbClient {
	return &lbClient{
		lb: lb(addrs),
	}
}

func (lb *lbClient) LBClient() *Client {
	return lb.lb.Client()
}
