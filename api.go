package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"domain"
)

func pokemonList() (*PokemonList, error) {
	url := "https://pokeapi.co/api/v2/pokemon"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making request: ", err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return nil, err
	}

	var pokemons PokemonList
	json.Unmarshal(responseBody, &pokemons)
	return &pokemons, nil

}

func pokemonData(url string)  {

}
