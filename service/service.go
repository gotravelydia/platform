// Copyright 2016 Travelydia, Inc. All rights reserved.

package service

import (
	"fmt"
	"net/http"

	"github.com/gotravelydia/platform/log"

	"github.com/gotravelydia/mux"
)

type Service struct {
	Name    string
	Version string
	Router  *mux.Router
}

func New(name, version string) *Service {
	log.Info(fmt.Sprintf("Initializing service: %s ver: %s", name, version))

	router := initRouter()
	service := &Service{
		Name:    name,
		Version: version,
		Router:  router,
	}

	return service
}

func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Handle invlaid routes.
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	// TODO: Add health monitor.
	router.HandleFunc("/", defaultPage)
	router.HandleFunc("/health", defaultPage)

	return router
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing to see here.")
}
