package repositories

import (
	"fmt"
	"pokeapi-go/src/models"
	"pokeapi-go/src/shared"
	"strings"
)

const apiUrl = "https://pokeapi.co/api/v2/"

type PokemonRepo interface {
	GetAllPokemons() (models.Resource, error)
	SearchPokemons(search string) (models.Resource, error)
	PokemonDetail(pokemonId string) (models.Pokemon, error)
}

type pokemonRepo struct {
}

func NewAddressRepo() PokemonRepo {
	return &pokemonRepo{}
}

func (r pokemonRepo) GetAllPokemons() (models.Resource, error) {
	endpoint := "pokemon?limit=20"

	var resource models.Resource

	shared.MakeRequest(apiUrl+endpoint, &resource)

	return resource, nil
}

func (r pokemonRepo) SearchPokemons(search string) (models.Resource, error) {
	endpoint := "pokemon?offset=0&limit=9999"

	var resource models.Resource

	shared.MakeRequest(apiUrl+endpoint, &resource)
	var newResults []models.Result

	for _, re := range resource.Results {

		if strings.Contains(re.Name, search) {
			newResults = append(newResults, re)
		}
	}

	resource.Results = newResults

	resource.Count = len(newResults)

	return resource, nil
}

func (r pokemonRepo) PokemonDetail(pokemonId string) (models.Pokemon, error) {

	endpoint := "pokemon/"
	var details models.Detail

	shared.MakeRequest(apiUrl+endpoint+pokemonId, &details)

	image := fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/home/%v.png", details.ID)

	var movesArr []string
	var typesArr []string

	for _, move := range details.Moves {
		movesArr = append(movesArr, move.Move.Name)
	}

	for _, _type := range details.Types {
		typesArr = append(typesArr, _type.Type.Name)
	}

	var stats []models.Stat

	for _, stat := range details.Stats {
		newStat := models.Stat{
			Name:  stat.Stat.Name,
			Value: stat.BaseStat,
		}

		stats = append(stats, newStat)
	}

	movesString := strings.Join(movesArr, ",")
	typesString := strings.Join(typesArr, ",")

	pokemon := models.Pokemon{
		ID:     details.ID,
		Name:   details.Name,
		Image:  image,
		Moves:  movesString,
		Type:   typesString,
		Weight: details.Weight,
		Height: details.Height,
		Stats:  stats,
	}

	return pokemon, nil

}
