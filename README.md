# microservice client
##使用方法
* 实现Service接口
<pre>type Service interface {
	GetServers() ([]Server, error)
}</pre>
* 初始化lb客户端,默认实现了随机访问均衡
<pre>
lbc := DefaultLBClient(&mockService{})
cli, err := lbc.LBClient().RestyClient()
req := cli.GetRequest()
if err != nil {
    t.Fatal(err)
	return
}
</pre>