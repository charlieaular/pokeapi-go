package services

import (
	"pokeapi-go/src/models"
	"pokeapi-go/src/repositories"
)

type PokemonService interface {
	GetAllPokemons() (models.Resource, error)
	SearchPokemons(search string) (models.Resource, error)
	PokemonDetail(pokemonId string) (models.Pokemon, error)
}

type pokemonService struct {
	repo repositories.PokemonRepo
}

func NewPokemonService(repo repositories.PokemonRepo) PokemonService {
	return &pokemonService{repo: repo}
}

func (s pokemonService) GetAllPokemons() (models.Resource, error) {
	return s.repo.GetAllPokemons()
}

func (s pokemonService) SearchPokemons(search string) (models.Resource, error) {
	return s.repo.SearchPokemons(search)
}

func (s pokemonService) PokemonDetail(pokemonId string) (models.Pokemon, error) {
	return s.repo.PokemonDetail(pokemonId)
}
