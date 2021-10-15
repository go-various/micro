package micro
// Least Connections
type lclb struct {
	Service Service
}
func (l *lclb)Client()*Client  {
	return nil
}