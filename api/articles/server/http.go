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
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) getArticleByIDRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: /id/:id", ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	vars := mux.Vars(r)
	id := vars["id"]

	serverSpan.SetTag("http.url", fmt.Sprintf("/id/%s", id))
	serverSpan.SetTag("http.method", "GET")

	artID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		serverSpan.SetTag("http.status_code", http.StatusBadRequest)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "could not parse parameter"}`, http.StatusBadRequest)
		return
	}

	article, err := httpsrv.dbm.GetArticleByID(uint64(artID))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON, err := json.Marshal(article)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	serverSpan.SetTag("http.status_code", 200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) getArticlesOfCategoryRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: /category/:category", ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	vars := mux.Vars(r)
	category := vars["category"]

	serverSpan.SetTag("http.url", fmt.Sprintf("/category/%s", category))
	serverSpan.SetTag("http.method", "GET")

	articles, err := httpsrv.dbm.GetArticlesOfCategory(category)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON, err := json.Marshal(articles)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	serverSpan.SetTag("http.status_code", 200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) getArticlesFromRegionRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: /region/:region", ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	vars := mux.Vars(r)
	regionIDStr := vars["region_id"]
	regionID, err := strconv.Atoi(regionIDStr)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "could not parse parameter"}`, http.StatusBadRequest)
		return
	}

	serverSpan.SetTag("http.url", fmt.Sprintf("/region/%s", regionIDStr))
	serverSpan.SetTag("http.method", "GET")

	articles, err := httpsrv.dbm.GetArticlesFromRegion(regionID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	respJSON, err := json.Marshal(articles)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	serverSpan.SetTag("http.status_code", 200)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) newArticleRoute(w http.ResponseWriter, r *http.Request) {
	var article db.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	newID, err := httpsrv.dbm.InsertArticle(article)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}
	article.ID = int64(newID)
	log.Println(article)

	respJSON, err := json.Marshal(article)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 201, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) newArticlesRoute(w http.ResponseWriter, r *http.Request) {
	var articles []db.Article
	err := json.NewDecoder(r.Body).Decode(&articles)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	var ids []int64

	for _, a := range articles {
		newID, err := httpsrv.dbm.InsertArticle(a)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
			return
		}
		ids = append(ids, int64(newID))
	}

	respJSON, err := json.Marshal(ids)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 201, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) assignRoutesToService() {
	httpsrv.r.HandleFunc("/", httpsrv.homeRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/id/{id}", httpsrv.getArticleByIDRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/category/{category}", httpsrv.getArticlesOfCategoryRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/region/id/{region_id}", httpsrv.getArticlesFromRegionRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/new", httpsrv.newArticleRoute).Methods(http.MethodPost)
	httpsrv.r.HandleFunc("/new/bulk", httpsrv.newArticlesRoute).Methods(http.MethodPost)
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
