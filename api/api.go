// Copyright 2016 Travelydia, Inc. All rights reserved.

package api

import (
	"net/http"

	"github.com/gotravelydia/platform/client/google"
	"github.com/gotravelydia/platform/config"
	"github.com/gotravelydia/platform/database"
	"github.com/gotravelydia/platform/log"
	"github.com/gotravelydia/platform/service"
)

const (
	AppServiceName       = "api"
	DefaultMaxFormMemory = 32 << 20 // 32 MB
)

type handleFunc func(w http.ResponseWriter, r *http.Request, ctx *Context)
type authLevel int

const (
	authLevelAnonymous    authLevel = iota // anonymous user with no uid, will require token in handler
	authLevelAnonymousUid                  // anonymous user with uid, checks session
	authLevelLogin                         // user login required, 401 if not available
	authLevelPayment                       // user payment verification required, 401 if not available
	authLevelAdmin                         // admin required, 401 if no user, 403 if not admin
	authLevelNone                          // no auth level for generic handler
)

type API struct {
	service          *service.Service
	db               *database.Database
	googleMapsClient *google.GoogleMaps
}

func New() (api *API, err error) {
	api = &API{}

	// Initialize API Router.
	api.initRouter()

	// Initialize the database connection.
	api.db, err = database.New()
	if err != nil {
		log.Error(err)
		return
	}

	// Initialize Google Maps API client.
	api.googleMapsClient, err = google.New()
	if err != nil {
		return
	}

	return
}

func (api *API) Close() {
	err := api.db.Close()
	if err != nil {
		log.Error(err)
	}
}

func (api *API) handler(h handleFunc, level authLevel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(DefaultMaxFormMemory)
		h(w, r, NewContext(api, r))
	}
}

func (api *API) initRouter() {
	api.service = service.New(
		AppServiceName,
		config.ServiceConfig.GetStringDefault("app.version", "0.0.0"),
	)

	searchRoute := api.service.Router.PathPrefix("/v1/search").Subrouter()
	searchRoute.HandleFunc("/location", api.handler(searchLocation, authLevelAnonymous)).Methods("GET")
}
