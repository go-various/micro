package micro

import "testing"

type mockHttpService struct {

}

func (s *mockHttpService)GetServers(name,tags string)([]Server,error){
	return []Server{
		{
			ID:          "moke-1",
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
type mockRPCService struct {

}

func (s *mockRPCService)GetServers(name, tags string)([]Server,error){
	return []Server{
		{
			ID:          "moke-1",
			Address:     "localhost:8080",
			Weight:      0,
			TPSDelay:    0,
			Connections: 0,
		},
	},nil
}

func TestClient_NewRPCCodecClient(t *testing.T) {
	lbc := DefaultLBClient(&mockRPCService{})

	cli, err := lbc.LBClient("","").NewRPCCodecClient()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	var rep string
	if err := cli.Call("RPC.Hello", "as", &rep); err != nil {
		t.Fatal(err)
		return
	}

}

func TestClient_NewRPCMsgpackClient(t *testing.T) {
	lbc := DefaultLBClient(&mockRPCService{})

	cli, err := lbc.LBClient("","").NewRPCMsgpackClient()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	var rep string

	if err := cli.Call("RPC.Hello", "hello", rep); err != nil {
		t.Fatal(err)
		return
	}

}
