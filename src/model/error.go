package model

import "net/http"

type Status struct {
	HttpCode int
	Code     int
	Message  string
}

var (
	Success    = Status{HttpCode: http.StatusOK, Code: 20000, Message: "success"}
	Error      = Status{HttpCode: http.StatusInternalServerError, Code: 50000, Message: "server error"}
	BadRequest = Status{HttpCode: http.StatusBadRequest, Code: 40000, Message: "bad request"}
)
