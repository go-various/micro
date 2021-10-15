
package micro

import (
	"errors"
	"math/rand"
)
// random
type rndlb struct {
	Service Service
}

func (l *rndlb) Client() *Client{
	if nil == l.Service {
		return &Client{serviceErr: errors.New("service not implemented")}
	}
	ss, err := l.Service.GetServers()
	if err != nil {
		return &Client{serviceErr: err}
	}
	s := ss[rand.Intn(len(ss))]
	return &Client{serviceErr: err, addr: s.Address}
}
