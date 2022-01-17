package micro

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type RestyClient struct {
	rawClient   *resty.Client
	enableTrace bool
	host        string
	hooks       []Hook
}

func DefaultResty(host string) *RestyClient {
	rawClient := resty.New()
	rawClient.SetTimeout(time.Second * 30)
	rawClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: false})
	rc := &RestyClient{
		rawClient:   rawClient,
		enableTrace: true,
		host:        host,
		hooks: make([]Hook, 0),
	}
	return rc
}

func NewResty(host string, timeout time.Duration, InsecureSkipVerify bool) *RestyClient {
	rawClient := resty.New()
	rawClient.SetTimeout(timeout)
	rawClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: InsecureSkipVerify})
	return &RestyClient{
		rawClient:   rawClient,
		enableTrace: true,
		host:        host,
		hooks: make([]Hook, 0),
	}
}

func (r *RestyClient) AddHooks(hooks... Hook)  {
	r.hooks = append(r.hooks, hooks...)
}

//GetRawClient 原生resty客户端
func (r *RestyClient) GetRawClient() *resty.Client {
	return r.rawClient
}

func (r *RestyClient) GetRequest() *request {
	req := r.rawClient.R()
	req.SetHeader("Span-ID", uuid.New().String())
	return &request{host: strings.TrimRight(r.host, "/"), Request: req, r: r}
}

//TraceInfo 调用日志
func (r *RestyClient) TraceInfo(resp *resty.Response, full bool) map[string]string {
	result := make(map[string]string)
	if full {
		ti := resp.Request.TraceInfo()
		result["DNSLookup"] = ti.DNSLookup.String()
		result["ConnTime"] = ti.ConnTime.String()
		result["TCPConnTime"] = ti.TCPConnTime.String()
		result["TLSHandshake"] = ti.TLSHandshake.String()
		result["ServerTime"] = ti.ServerTime.String()
		result["ResponseTime"] = ti.ResponseTime.String()
		result["TotalTime"] = ti.TotalTime.String()
		result["IsConnReused"] = strconv.FormatBool(ti.IsConnReused)
		result["IsConnWasIdle"] = strconv.FormatBool(ti.IsConnWasIdle)
		result["ConnIdleTime"] = ti.ConnIdleTime.String()
	} else {
		result["Response Info"] = "Response"
		result["code"] = strconv.Itoa(resp.StatusCode())
		result["status"] = resp.Status()
		result["proto"] = resp.Proto()
		result["time"] = resp.Time().String()
		result["received_time"] = resp.ReceivedAt().String()
		result["body"] = fmt.Sprintf("%s", resp)
	}
	return result
}
