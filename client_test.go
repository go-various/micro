package micro

import "testing"

func TestDefaultRestyClient(t *testing.T) {
	lbc := DefaultClient(map[string][]Server{
		"http://127.0.0.1/jdsh-v1": []Server{
			{
			},
		},
		"http://127.0.0.2/jdsh-v1": []Server{
			{
			},
		},
		"http://127.0.0.3/jdsh-v1": []Server{
			{
			},
		},
	})
	cli := lbc.LBClient()

	res, err := cli.RestyClient().GetRequest().Post("/names")
	t.Log(err, res)
}

func TestClient_NewRPCCodecClient(t *testing.T) {
	lbc := DefaultClient(map[string][]Server{
		"127.0.0.1:7109": []Server{
			{
			},
		},
	})

	cli, err := lbc.LBClient().NewRPCCodecClient()
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
	lbc := DefaultClient(map[string][]Server{
		"127.0.0.1:7109": []Server{
			{
			},
		},
	})

	cli, err := lbc.LBClient().NewRPCMsgpackClient()
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
