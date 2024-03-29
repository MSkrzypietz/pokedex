package main

import (
	"github.com/MSkrzypietz/pokedex/pokeapi"
	"time"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasUrl *string
	prevLocationAreasUrl *string
	caughtPokemon        map[string]pokeapi.PokemonResp
}

func main() {
	cfg := config{
		pokeapiClient:        pokeapi.NewClient(time.Hour),
		nextLocationAreasUrl: nil,
		prevLocationAreasUrl: nil,
		caughtPokemon:        make(map[string]pokeapi.PokemonResp),
	}
	startRepl(&cfg)
}
