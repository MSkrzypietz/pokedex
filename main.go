package main

import "github.com/MSkrzypietz/pokedex/pokeapi"

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasUrl *string
	prevLocationAreasUrl *string
}

func main() {
	cfg := config{
		pokeapiClient:        pokeapi.NewClient(),
		nextLocationAreasUrl: nil,
		prevLocationAreasUrl: nil,
	}
	startRepl(&cfg)
}
