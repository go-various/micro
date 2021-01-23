// Consistent Hash
package micro


type hashlb struct {
	addrs Servers
}

func (l *hashlb)Client()*Client  {
	return nil
}