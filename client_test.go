package micro

import (
	"net/http"
	"testing"
)

type mockHttpService struct {
}

type trace struct {
}

func (t trace) Trace(req *http.Request, res *http.Response, err error) {
}

func (s *mockHttpService) GetServers(name, tags string) ([]Server, error) {
	return []Server{}, nil
}

func TestDefaultRestyClient(t *testing.T) {
	lbc := RandomAdapter(&mockHttpService{})
	lbc.AddHooks(&trace{})
	cli := lbc.Client("mock-1", "")

	rc, err := cli.RestyClient()
	if err != nil {
		t.Fatal(err)
	}
	req := rc.GetRequest()
	res, err := req.Post("/names")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(res)
}

type mockRPCService struct {
}

func (s *mockRPCService) GetServers(name, tags string) ([]Server, error) {
	return []Server{
		{
			ID:          "mock-1",
			Address:     "localhost:8080",
			Weight:      0,
			TPSDelay:    0,
			Connections: 0,
		},
	}, nil
}

func TestClient_NewRPCCodecClient(t *testing.T) {
	lbc := RandomAdapter(&mockRPCService{})

	cli, err := lbc.Client("", "").NewRPCCodecClient()
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
	lbc := RandomAdapter(&mockRPCService{})

	cli, err := lbc.Client("", "").NewRPCMsgpackClient()
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
