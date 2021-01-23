package micro

import "errors"

type Policy string

const (
	PolicyRandom           Policy = "lb:rnd"
	PolicyRoundRobin       Policy = "lb:rr"
	PolicyConsistentHash   Policy = "lb:hash"
	PolicyLeastConnections Policy = "lb:lcs"
)

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

type Servers map[string][]Server

type lbClient struct {
	lb LoadBalance
}

func DefaultClient(addrs Servers) *lbClient {
	return &lbClient{
		lb: &rndlb{addrs: addrs},
	}
}

func NewClient(p Policy,addrs Servers) (*lbClient,error) {
	switch p {
	case PolicyRandom:
		return &lbClient{lb: &rndlb{addrs: addrs}},nil

	case PolicyRoundRobin:
		return &lbClient{lb: &rrlb{addrs: addrs}},nil

	case PolicyConsistentHash:
		return &lbClient{lb: &hashlb{addrs: Servers{}}},nil

	case PolicyLeastConnections:
		return &lbClient{lb: &lclb{addrs: Servers{}}}, nil

	default:
		return nil, errors.New("invalid loadbalance policy")
	}
}

func (lb *lbClient) LBClient() *Client {
	return lb.lb.Client()
}
