// Consistent Hash
package micro


type hashlb struct {
	addrs Addrs
}

func (l *hashlb)Client()*Client  {
	return nil
}