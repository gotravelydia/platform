// Copyright 2016 Travelydia, Inc. All rights reserved.

package api

import (
	"net/http"
)

type Context struct {
	*API
	errorTags map[string]string
}

func NewContext(api *API, r *http.Request) *Context {
	ctx := &Context{
		API:       api,
		errorTags: make(map[string]string),
	}

	return ctx
}
