
package micro

// Consistent Hash
type hashlb struct {
	Service Service
}

func (l *hashlb)Client()*Client  {
	return nil
}