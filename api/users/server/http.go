package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type httpService struct {
	r *mux.Router
	// TODO: add copy of mysql client
}

func (httpsrv *httpService) homeRoute(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	ipList, _ := net.LookupHost(host)
	var ip string = "0.0.0.0"
	if len(ipList) > 0 {
		ip = ipList[0]
	}

	// create response
	type IPData struct {
		IP   string `json:"ip"`
		Host string `json:"server"`
	}
	ipData := IPData{IP: ip, Host: host}
	ipDataJSON, err := json.Marshal(ipData)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(ipDataJSON))
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
