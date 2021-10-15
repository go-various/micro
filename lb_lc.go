package micro
// Least Connections
type lclb struct {
	Service Service
}
func (l *lclb)Client(name, tags string)*Client  {
	return nil
}