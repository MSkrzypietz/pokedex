package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	resp, err := c.httpClient.Get(apiUrl)
	if err != nil {
		return locationAreas, errors.New("unable to fetch the location areas")
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreas, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreas, errors.New("couldn't read request")
	}

	if err = json.Unmarshal(body, &locationAreas); err != nil {
		return locationAreas, errors.New("couldn't unmarshal parameters")
	}
	return locationAreas, nil
}
