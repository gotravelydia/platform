// Copyright 2016 Travelydia, Inc. All rights reserved.

package service

import (
	"errors"
	"net/http"

	"github.com/gotravelydia/platform/log"

	"github.com/bitly/go-simplejson"
)

func SetJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func Success200(w http.ResponseWriter, payload []byte) {
	SetJsonContentType(w)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func Success200Empty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotFound, errors.New("API not found for "+r.URL.Path))
}

func writeError(w http.ResponseWriter, code int, err error) {
	SetJsonContentType(w)
	w.WriteHeader(code)
	w.Write(errorResponse(code, err))
}

func errorResponse(code int, err error) []byte {
	json := simplejson.New()
	json.Set("code", code)
	if err != nil {
		json.Set("error", err.Error())
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Error(err)
		return nil
	}

	return payload
}
