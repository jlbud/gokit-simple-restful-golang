package main

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

var (
	res, a, b int
	calError  error
)

// 逻辑层
// 注意 type endpoint.Endpoint
func MakeEndpoint(ser Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		fmt.Println("this is MakeEndpoint")
		req := request.(ArithmeticRequest)
		switch req.RequestType {
		case "Add":
			a = req.A
			b = req.B
			res = ser.Add(a, b)
		default:
			res = 100
		}
		return ArithmeticResponse{Result: res, Error: nil}, nil
	}
}
