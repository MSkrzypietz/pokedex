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
