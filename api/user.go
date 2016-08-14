// Copyright 2016 Travelydia, Inc. All rights reserved.

package api

import (
	"net/http"

	"github.com/gotravelydia/platform/log"
	"github.com/gotravelydia/platform/service"
)

// Returns a new token for the anonymous user.
func getNewToken(w http.ResponseWriter, r *http.Request) {
	log.Error("Hello world!")

	// Return with no user payload.
	service.Success200Empty(w)
}
