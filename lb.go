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

type LBClient struct {
	lb LoadBalance
}

func DefaultLBClient(service Service) *LBClient {
	return &LBClient{
		lb: &rndlb{Service: service},
	}
}

func NewLBClient(p Policy, service Service) *LBClient {
	switch p {
	case PolicyRandom:
		return &LBClient{lb: &rndlb{Service: service}}

	case PolicyRoundRobin:
		return &LBClient{lb: &rrlb{Service: service}}

	case PolicyConsistentHash:
		return &LBClient{lb: &hashlb{Service: service}}

	case PolicyLeastConnections:
		return &LBClient{lb: &lclb{Service: service}}

	default:
		return nil
	}
}

func (lb *LBClient) LBClient(name, tags string) *Client {
	return lb.lb.Client(name, tags)
}
