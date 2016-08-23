// Copyright 2016 Travelydia, Inc. All rights reserved.

package api

import (
	"errors"
	"net/http"

	"github.com/gotravelydia/platform/model"
	"github.com/gotravelydia/platform/service"
)

func searchLocation(w http.ResponseWriter, r *http.Request, ctx *Context) {
	q := r.FormValue(model.SearchQuery)

	locations, err := ctx.googleMapsClient.SearchPlaces(q)
	if err != nil {
		service.InternalServerError(w, errors.New("Failed to get search results."))
		return
	}

	payload, err := model.BuildLocationArrayPayload(locations)
	if err != nil {
		service.InternalServerError(w, errors.New("Failed to build search results."))
		return
	}

	service.OK(w, payload)
}
