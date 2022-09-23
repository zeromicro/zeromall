package router

import (
	"net/http"

	"github.com/better-go/pkg/log"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ServiceFunc func(r *http.Request) (resp interface{}, err error)

func HandlerWrap(fn ServiceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// biz logic handle:
		resp, err := fn(r)
		log.Infof("http api: url=%+v, req=%+v, resp=%+v, err=%v", r.RequestURI, r.Form, resp, err)

		// resp:
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
