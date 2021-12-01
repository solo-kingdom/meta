package main

import (
	"fmt"
	"github.com/solo-kingdom/meta/pkg/settings"
	"github.com/solo-kingdom/meta/src/router"
	"net/http"
)

func init() {
	settings.SetUp()
}

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.ServerConfig.Port),
		Handler: router.GetRouter(),
	}

	err := server.ListenAndServe()
	if err != nil {
		println(err.Error())
	}
}
