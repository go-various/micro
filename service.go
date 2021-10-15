package micro


//Service is 服务列表获取接口

type Service interface {
	//GetServers 获取服务列表
	//name 微服务名字
	//tags 逗号分割的标签
	GetServers(name string, tags string) ([]Server, error)
}