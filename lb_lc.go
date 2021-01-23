// Least Connections
package micro

type lclb struct {
	addrs Addrs
}

func (l *lclb)Client()*Client  {
	return nil
}