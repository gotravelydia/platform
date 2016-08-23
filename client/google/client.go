// Copyright 2016 Travelydia, Inc. All rights reserved.

package google

import (
	"github.com/gotravelydia/platform/config"
	"github.com/gotravelydia/platform/log"
	"github.com/gotravelydia/platform/model"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

type GoogleMaps struct {
	Client *maps.Client
}

func New() (*GoogleMaps, error) {
	apiKey, err := config.ServiceConfig.GetString("google.maps.api_key")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &GoogleMaps{
		Client: c,
	}, nil
}

func (c *GoogleMaps) SearchPlaces(q string) ([]*model.Location, error) {
	r := &maps.PlaceAutocompleteRequest{
		Input: q,
		Type:  "geocode",
	}

	resp, err := c.Client.PlaceAutocomplete(context.Background(), r)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	locations := []*model.Location{}
	for _, res := range resp.Predictions {
		location := &model.Location{
			PlaceId:     res.PlaceID,
			Description: res.Description,
		}

		locations = append(locations, location)
	}

	return locations, nil
}
