package micro

// round robin


type rrlb struct {
	Service Service
}

func (l *rrlb) Client() *Client {
	return nil
}