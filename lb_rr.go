package micro

// round robin


type rrlb struct {
	addrs Servers
}

func (l *rrlb) Client() *Client {
	return nil
}