package model

import "net/http"

type Status struct {
	HttpCode int
	Success  bool
	Code     int
	Message  string
}

var (
	Success    = Status{HttpCode: http.StatusOK, Success: true, Code: 20000, Message: "success"}
	Error      = Status{HttpCode: http.StatusInternalServerError, Code: 50000, Message: "server e"}
	BadRequest = Status{HttpCode: http.StatusBadRequest, Code: 40000, Message: "bad request"}
	NotFound   = Status{HttpCode: http.StatusBadRequest, Code: 40400, Message: "resource not found"}
)
