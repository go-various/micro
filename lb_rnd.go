// random
package micro

import (
	"math/rand"
)

type rndlb struct {
	addrs Servers
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
