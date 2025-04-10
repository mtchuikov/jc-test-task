package v1handlers

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type prepareResponseArgs[RespT any] struct {
	statusCode int
	response   *RespT
}

func prepareResponse[RespT any](rw http.ResponseWriter, resp *prepareResponseArgs[RespT]) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.statusCode)
	payload, _ := jsoniter.Marshal(resp.response)
	rw.Write(payload)
}
