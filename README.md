# microservice client
##使用方法
* 实现Service接口
<pre>type Service interface {
	GetServers() ([]Server, error)
}</pre>
see [client_test.go](client_test.go)