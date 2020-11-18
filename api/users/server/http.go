package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frouioui/tagenal/api/users/pb"
	"github.com/gorilla/mux"
)

type httpService struct {
	r *mux.Router
	// TODO: add copy of mysql client
}

func (httpsrv *httpService) homeRoute(w http.ResponseWriter, r *http.Request) {
	resp := &pb.UserHomeResponse{
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

func newServiceHTTP() (httpsrv httpService) {
	httpsrv.r = mux.NewRouter()
	httpsrv.assignRoutesToService()
	return httpsrv
}
