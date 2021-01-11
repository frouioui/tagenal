package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/frouioui/tagenal/api/users/db"
	"github.com/frouioui/tagenal/api/users/pb"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type httpService struct {
	r   *mux.Router
	dbm *db.DatabaseManager
}

func (httpsrv *httpService) homeRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: /", ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	resp := &pb.InformationResponse{
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

func (httpsrv *httpService) healthRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: "+r.RequestURI, ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	w.WriteHeader(200)
	fmt.Fprint(w, `ok`)
}

func (httpsrv *httpService) readyRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: "+r.RequestURI, ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	if ready == false {
		w.WriteHeader(500)
		fmt.Fprint(w, `ko`)
		return
	}
	w.WriteHeader(200)
	fmt.Fprint(w, `ok`)
}

func (httpsrv *httpService) getUserByIDRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: /id/:id", ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	vars := mux.Vars(r)
	id := vars["id"]

	serverSpan.SetTag("http.url", fmt.Sprintf("/id/%s", id))
	serverSpan.SetTag("http.method", "GET")

	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		serverSpan.SetTag("http.status_code", http.StatusBadRequest)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "could not parse parameter"}`, http.StatusBadRequest)
		return
	}

	vtspanctx, err := getVitessSpanContextFromTextMap(serverSpan.Context())
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	var user db.User
	if user, err = getCacheUser(opentracing.ContextWithSpan(context.Background(), serverSpan), fmt.Sprintf("user_id_%d", userID), user); err != nil {
		user, err = httpsrv.dbm.GetUserByID(opentracing.ContextWithSpan(context.Background(), serverSpan), vtspanctx, uint64(userID))
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			serverSpan.SetTag("http.status_code", http.StatusInternalServerError)
			fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
			return
		}
		err = setCacheUser(opentracing.ContextWithSpan(context.Background(), serverSpan), fmt.Sprintf("user_id_%d", userID), user)
		if err != nil {
			log.Println(err.Error())
		}
	}

	respJSON, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		serverSpan.SetTag("http.status_code", http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(http.StatusOK)
	serverSpan.SetTag("http.status_code", http.StatusOK)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) getUsersOfRegionRoute(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	serverSpan := tracer.StartSpan("HTTP GET URL: /region/:region", ext.RPCServerOption(spanCtx))
	defer serverSpan.Finish()

	vars := mux.Vars(r)
	region := vars["region"]

	serverSpan.SetTag("http.url", fmt.Sprintf("/region/%s", region))
	serverSpan.SetTag("http.method", "GET")

	vtspanctx, err := getVitessSpanContextFromTextMap(serverSpan.Context())
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		serverSpan.SetTag("http.status_code", 500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	var users []db.User
	if users, err = getCacheUsers(opentracing.ContextWithSpan(context.Background(), serverSpan), fmt.Sprintf("user_region_%s", region), users); err != nil {
		users, err = httpsrv.dbm.GetUsersOfRegion(opentracing.ContextWithSpan(context.Background(), serverSpan), vtspanctx, region)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			serverSpan.SetTag("http.status_code", http.StatusInternalServerError)
			fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
			return
		}
		err = setCacheUsers(opentracing.ContextWithSpan(context.Background(), serverSpan), fmt.Sprintf("user_region_%s", region), users)
		if err != nil {
			log.Println(err.Error())
		}
	}

	respJSON, err := json.Marshal(users)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		serverSpan.SetTag("http.status_code", http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(http.StatusOK)
	serverSpan.SetTag("http.status_code", http.StatusOK)
	fmt.Fprintf(w, `{"status": "success", "code": 200, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) newUserRoute(w http.ResponseWriter, r *http.Request) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	newID, err := httpsrv.dbm.InsertUser(user)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = int64(newID)

	respJSON, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, "server error")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"status": "success", "code": 201, "data": %s}`, string(respJSON))
}

func (httpsrv *httpService) newUsersRoute(w http.ResponseWriter, r *http.Request) {
	var users []db.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"status": "failure", "code": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())
		return
	}

	var ids []int64

	for _, u := range users {
		newID, err := httpsrv.dbm.InsertUser(u)
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
	httpsrv.r.HandleFunc("/health", httpsrv.healthRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/ready", httpsrv.readyRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/id/{id}", httpsrv.getUserByIDRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/region/{region}", httpsrv.getUsersOfRegionRoute).Methods(http.MethodGet)
	httpsrv.r.HandleFunc("/new", httpsrv.newUserRoute).Methods(http.MethodPost)
	httpsrv.r.HandleFunc("/new/bulk", httpsrv.newUsersRoute).Methods(http.MethodPost)
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
