// Copyright 2016 Travelydia, Inc. All rights reserved.

package api

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gotravelydia/platform/config"
	"github.com/gotravelydia/platform/log"
)

func Serve() {
	api, err := New()
	if err != nil {
		log.Fatal(err)
	}
	defer api.Close()

	runtime.GOMAXPROCS((runtime.NumCPU() * 2) + 1)

	port := config.ServiceConfig.GetIntDefault("app.port", 8100)
	http.ListenAndServe(fmt.Sprintf(":%d", port), api.service.Router)
}
