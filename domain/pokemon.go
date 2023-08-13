package domain

type PokemonList struct {
	Count int
	Next string
	Previous string
	Results []Pokemon
}

type Pokemon struct {
	Name string 
	Url string
}