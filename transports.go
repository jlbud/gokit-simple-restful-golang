package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

// 把请求参数解析为可用参数
func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	l := len(vars)
	if l == 0 {
		return ArithmeticRequest{}, nil
	}

	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}
	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequest
	}
	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}
	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)

	return ArithmeticRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
}

// 响应结果餐素
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// 请求路由层
func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	router.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint,
		decodeRequest,
		encodeResponse,
		options...,
	))

	router.Methods("GET").Path("/calculate/result").Handler(kithttp.NewServer(
		endpoint,
		decodeRequest,
		encodeResponse,
		options...,
	))
	return router
}
