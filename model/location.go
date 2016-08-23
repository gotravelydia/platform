// Copyright 2016 Travelydia, Inc. All rights reserved.

package model

import (
	"net/http"
	"strconv"

	"github.com/gotravelydia/platform/log"

	"github.com/bitly/go-simplejson"
)

type Location struct {
	Id          int64  `json:"id"`
	PlaceId     string `json:"placeId"`
	Description string `json:"description"`
}

func GetLocationFromRequest(r *http.Request) (*Location, error) {
	var err error

	var locationId int64
	locationIdStr := r.FormValue("locationId")
	if locationIdStr != "" {
		locationId, err = strconv.ParseInt(locationIdStr, 0, 64)
		if err != nil {
			return nil, err
		}
	}

	return &Location{
		Id:          locationId,
		PlaceId:     r.FormValue("placeId"),
		Description: r.FormValue("description"),
	}, nil
}

func BuildLocationArrayPayload(locations []*Location) ([]byte, error) {
	payload, err := buildLocationArray(locations).MarshalJSON()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return payload, nil
}

func buildLocationArray(locations []*Location) *simplejson.Json {
	json := simplejson.New()

	jsonArray, _ := json.Array()
	for _, location := range locations {
		jsonArray = append(jsonArray, buildLocationJsonObject(location))
	}

	json.Set("locations", jsonArray)

	return json
}

func buildLocationJsonObject(location *Location) *simplejson.Json {
	locationJson := simplejson.New()
	locationJson.Set("id", location.Id)
	locationJson.Set("placeId", location.PlaceId)
	locationJson.Set("description", location.Description)
	return locationJson
}
