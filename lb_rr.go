package micro

// round robin


type rrlb struct {
	Service Service
}

func (l *rrlb) Client(name, tags string) *Client {
	return nil
}