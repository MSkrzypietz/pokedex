package pokeapi

import (
	"encoding/json"
	"errors"
)

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	apiUrl := baseApiUrl + "location-area"
	if pageUrl != nil {
		apiUrl = *pageUrl
	}

	var locationAreas LocationAreasResp
	data, err := c.httpGet(apiUrl)
	if err != nil {
		return locationAreas, err
	}

	if err = json.Unmarshal(data, &locationAreas); err != nil {
		return locationAreas, errors.New("couldn't unmarshal parameters")
	}
	return locationAreas, nil
}
