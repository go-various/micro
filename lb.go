package micro

type Policy string

const (
	PolicyRandom           Policy = "lb:rnd"
	PolicyRoundRobin       Policy = "lb:rr"
	PolicyConsistentHash   Policy = "lb:hash"
	PolicyLeastConnections Policy = "lb:lcs"
)

type LoadBalance interface {
	Client(name, tags string) *Client
}

type Server struct {
	ID          string
	Address     string
	Weight      int
	TPSDelay    float64
	Connections int
}

type Servers map[string][]Server

type lbClient struct {
	lb LoadBalance
}

func DefaultLBClient(service Service) *lbClient {
	return &lbClient{
		lb: &rndlb{Service: service},
	}
}

func NewLBClient(p Policy, service Service) *lbClient {
	switch p {
	case PolicyRandom:
		return &lbClient{lb: &rndlb{Service: service}}

	case PolicyRoundRobin:
		return &lbClient{lb: &rrlb{Service: service}}

	case PolicyConsistentHash:
		return &lbClient{lb: &hashlb{Service: service}}

	case PolicyLeastConnections:
		return &lbClient{lb: &lclb{Service: service}}

	default:
		return nil
	}
}

func (lb *lbClient) LBClient(name, tags string) *Client {
	return lb.lb.Client(name, tags)
}
