package micro

import "testing"

func TestDefaultRestyClient(t *testing.T) {
	lbc := NewLBClient(map[string][]Meta{
		"http://127.0.0.1/jdsh-v1": []Meta{
			{
				Name:  "weight",
				Value: "123",
			},
		},
		"http://127.0.0.2/jdsh-v1": []Meta{
			{
				Name:  "weight",
				Value: "123",
			},
		},
		"http://127.0.0.3/jdsh-v1": []Meta{
			{
				Name:  "weight",
				Value: "123",
			},
		},
	})
	cli := lbc.LBClient()

	res, err := cli.DefaultRestyClient().GetRequest().Post("/names")
	t.Log(err, res)
}

func TestClient_NewRPCCodecClient(t *testing.T) {
	lbc := NewLBClient(map[string][]Meta{
		"127.0.0.1:7109": []Meta{
			{
				Name:  "weight",
				Value: "123",
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
	lbc := NewLBClient(map[string][]Meta{
		"127.0.0.1:7109": []Meta{
			{
				Name:  "weight",
				Value: "123",
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