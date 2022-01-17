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

type LBAdapter interface {
	Client(name, tags string) *Client
	AddHooks(hooks ...Hook)
}

type lbClient struct {
	lb    LoadBalance
	hooks []Hook
}

func (lbc *lbClient) Client(name, tags string) *Client {
	cli := lbc.lb.Client(name, tags)
	cli.hooks = lbc.hooks
	return cli
}

func (lbc *lbClient) AddHooks(hooks ...Hook) {
	lbc.hooks = append(lbc.hooks, hooks...)
}
