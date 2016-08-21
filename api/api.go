// Copyright 2016 Travelydia, Inc. All rights reserved.

package api

import (
	"net/http"

	"github.com/gotravelydia/platform/config"
	"github.com/gotravelydia/platform/database"
	"github.com/gotravelydia/platform/log"
	"github.com/gotravelydia/platform/service"
)

const (
	AppServiceName       = "api"
	DefaultMaxFormMemory = 32 << 20 // 32 MB
)

type handleFunc func(w http.ResponseWriter, r *http.Request)
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
	service *service.Service
	db      *database.Database
}

func New() (*API, error) {
	api := &API{}

	// Initialize API Router.
	api.initRouter()

	// Initialize the database connection.
	db, err := database.New()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	api.db = db

	return api, nil
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
		h(w, r)
	}
}

func (api *API) initRouter() {
	api.service = service.New(
		AppServiceName,
		config.ServiceConfig.GetStringDefault("app.version", "0.0.0"),
	)

	userRoute := api.service.Router.PathPrefix("/v1/users").Subrouter()
	userRoute.HandleFunc("/hello", api.handler(getNewToken, authLevelAnonymous)).Methods("GET", "POST")
}
