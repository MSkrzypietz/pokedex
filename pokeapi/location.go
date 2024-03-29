package pokeapi

import (
	"encoding/json"
	"errors"
)

type LocationResp struct {
	ID        int    `json:"id"`
	GameIndex int    `json:"game_index"`
	Name      string `json:"name"`
	Location  struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocation(name string) (LocationResp, error) {
	apiUrl := baseApiUrl + "location-area/" + name

	var locationArea LocationResp
	data, err := c.httpGet(apiUrl)
	if err != nil {
		return locationArea, err
	}

	if err = json.Unmarshal(data, &locationArea); err != nil {
		return locationArea, errors.New("couldn't unmarshal parameters")
	}
	return locationArea, nil
}
