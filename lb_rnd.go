//+build lb_rnd

// random
package micro

import (
	"math/rand"
)

func init() {
	lb = func(addrs Addrs) LoadBalance {
		return &rndlb{addrs: addrs}
	}
}

var lb Factory

type rndlb struct {
	addrs Addrs
}

func (l *rndlb) Client() *Client {

	var addrs []string
	for addr, _ := range l.addrs {
		addrs = append(addrs, addr)
	}
	if len(addrs) == 0 {
		return nil
	}
	return &Client{addr: addrs[rand.Intn(len(addrs))]}
}
