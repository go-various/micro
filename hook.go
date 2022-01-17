package micro

import "net/http"

type Hook interface {
	Trace(*http.Request, *http.Response, error)
}