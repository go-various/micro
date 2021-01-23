package micro

// round robin


type rrlb struct {
	addrs Addrs
}

func (l *rrlb) Client() *Client {
	return nil
}