package micro

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type request struct {
	host string
	*resty.Request
	r *RestyClient
}

func (r *request) SetQueryParam(param, value string) *request {
	r.Request.QueryParam.Set(param, value)
	return r
}

func (r *request) SetQueryParams(params map[string]string) *request {
	r.Request.SetQueryParams(params)
	return r
}

func (r *request) SetQueryParamsFromValues(params url.Values) *request {
	r.Request.SetQueryParamsFromValues(params)
	return r
}
func (r *request) SetFormData(data map[string]string) *request {
	r.Request.SetFormData(data)
	return r
}

func (r *request) SetFormDataFromValues(data url.Values) *request {
	r.Request.SetFormDataFromValues(data)
	return r
}

func (r *request) SetMultipartFormData(data map[string]string) *request {
	r.Request.SetMultipartFormData(data)
	return r
}

func (r *request) SetFile(param, filePath string) *request {
	r.Request.SetFile(param, filePath)
	return r
}
func (r *request) SetFiles(files map[string]string) *request {
	r.Request.SetFiles(files)
	return r
}
func (r *request) SetFileReader(param, fileName string, reader io.Reader) *request {
	r.Request.SetFileReader(param, fileName, reader)
	return r
}
func (r *request) SetMultipartField(param, fileName, contentType string, reader io.Reader) *request {
	r.Request.SetMultipartField(param, fileName, contentType, reader)
	return r
}

func (r *request) SetBody(body interface{}) *request {
	r.Request.SetBody(body)
	return r
}

func (r *request) SetHeader(header, value string) *request {
	r.Request.SetHeader(header, value)
	return r
}
func (r *request) SetHeaders(headers map[string]string) *request {
	r.Request.SetHeaders(headers)
	return r
}

func (r *request) Post(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodPost, path)
}

func (r *request) Get(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodGet, path)
}

func (r *request) Put(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodPut, path)
}

func (r *request) Delete(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodDelete, path)
}

func (r *request) Patch(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodPatch, path)
}

func (r *request) Head(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodHead, path)
}

func (r *request) Options(path string) (*resty.Response, error) {
	return r.Execute(resty.MethodOptions, path)
}

func (r *request) Execute(method string, path string) (resp *resty.Response, err error) {
	uri := fmt.Sprintf("%s/%s", r.host, strings.TrimLeft(path, "/"))
	if circuit.IsHolding(uri) {
		r.trace(r.RawRequest, nil, ErrCircuitBreakerMessage)
		return nil, ErrCircuitBreakerMessage
	}

	resp, err = r.Request.Execute(method, uri)
	if err != nil {
		circuit.Failed(uri)
		return nil, err
	}

	r.trace(r.RawRequest, resp.RawResponse, nil)
	return resp, nil

}

func (r *request) trace(req *http.Request, res *http.Response, err error) {
	for _, hook := range r.r.hooks {
		hook.Trace(req, res, err)
	}
}
