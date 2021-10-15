
package micro

// Consistent Hash
type hashlb struct {
	Service Service
}

func (l *hashlb)Client(name, tags string)*Client  {
	return nil
}