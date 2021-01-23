// Least Connections
package micro

type lclb struct {
	addrs Servers
}

func (l *lclb)Client()*Client  {
	return nil
}