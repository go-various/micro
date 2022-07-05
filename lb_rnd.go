
package micro

import (
	"errors"
	"math/rand"
)
// random loadBalance client
type rndlb struct {
	Service Service
}

func RandomAdapter(service Service) ClientAdapter {
	return &lbClient{
		lb:    &rndlb{Service: service},
		hooks: make([]Hook, 0),
	}
}

func (l *rndlb) Client(name string, tags string) *Client{
	if nil == l.Service {
		return &Client{serviceErr: errors.New("service not implemented")}
	}
	ss, err := l.Service.GetServers(name, tags)
	if err != nil {
		return &Client{serviceErr: err}
	}
	if len(ss) == 0 {
		return &Client{serviceErr: errors.New("services length is zero")}
	}

	s := ss[rand.Intn(len(ss))]
	return &Client{serviceErr: err, addr: s.Address}
}
