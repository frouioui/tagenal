package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frouioui/tagenal/api/articles/pb"
	"github.com/gorilla/mux"
)

type httpService struct {
	r *mux.Router
}

func (httpsrv *httpService) homeRoute(w http.ResponseWriter, r *http.Request) {
	resp := &pb.ArticleHomeResponse{
		IP:   getHostIP(),
		Host: getHostName(),
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) assignRoutesToService() {
	httpsrv.r.HandleFunc("/", httpsrv.homeRoute).Methods(http.MethodGet)
}

func (httpsrv *httpService) getRouter() *mux.Router {
	return httpsrv.r
}

func newServiceHTTP() (httpsrv httpService, err error) {
	httpsrv.r = mux.NewRouter()
	httpsrv.assignRoutesToService()
	return httpsrv, nil
}
