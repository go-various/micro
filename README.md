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
res, err := req.Post("/names")
t.Log(err, res)
</pre>
*example
<pre>
type mockHttpService struct {

}

func (s *mockHttpService)GetServers(name,tags string)([]Server,error){
	return []Server{
		{
			ID:          "mock-1",
			Address:     "http://localhost:8080",
			Weight:      0,
			TPSDelay:    0,
			Connections: 0,
		},
	},nil
}

func TestDefaultRestyClient(t *testing.T) {
	lbc := DefaultLBClient(&mockHttpService{})

	cli, err := lbc.LBClient("","").RestyClient()
	req := cli.GetRequest()
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := req.Post("/names")
	t.Log(err, res)
}
</pre>