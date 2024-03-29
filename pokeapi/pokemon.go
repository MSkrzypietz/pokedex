package pokeapi

import (
	"encoding/json"
	"errors"
)

type PokemonResp struct {
	ID             int    `json:"id"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (PokemonResp, error) {
	apiUrl := baseApiUrl + "pokemon/" + name

	var pokemonResp PokemonResp
	data, err := c.httpGet(apiUrl)
	if err != nil {
		return pokemonResp, err
	}

	if err = json.Unmarshal(data, &pokemonResp); err != nil {
		return pokemonResp, errors.New("couldn't unmarshal parameters")
	}
	return pokemonResp, nil
}
