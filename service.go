package micro

//Service is 服务列表获取接口
type Service interface {
	GetServers(name string) ([]Server, error)
}