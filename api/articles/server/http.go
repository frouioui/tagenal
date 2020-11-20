package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/frouioui/tagenal/api/articles/db"
	"github.com/frouioui/tagenal/api/articles/pb"
	"github.com/gorilla/mux"
)

type httpService struct {
	r   *mux.Router
	dbm *db.DatabaseManager
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

func (httpsrv *httpService) getArticleByIDRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	artID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "could not parse parameter"}`, http.StatusBadRequest)
		return
	}

	article, err := httpsrv.dbm.GetArticleByID(uint64(artID))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) getArticlesOfCategoryRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	region := vars["region"]

	articles, err := httpsrv.dbm.GetArticlesOfRegion(region)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON, err := json.Marshal(articles)
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
	httpsrv.r.HandleFunc("/id/{id}", httpsrv.getArticleByIDRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/category/{category}", httpsrv.getArticlesOfCategoryRoute).Methods(http.MethodGet)
}

func (httpsrv *httpService) getRouter() *mux.Router {
	return httpsrv.r
}

func newServiceHTTP() (httpsrv httpService, err error) {
	httpsrv.dbm, err = db.NewDatabaseManager()
	if err != nil {
		log.Println(err.Error())
		return httpsrv, err
	}
	httpsrv.r = mux.NewRouter()
	httpsrv.assignRoutesToService()
	return httpsrv, nil
}
