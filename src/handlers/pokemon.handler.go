package handlers

import (
	"net/http"
	"pokeapi-go/src/models"
	"pokeapi-go/src/services"
	"pokeapi-go/src/shared"

	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	pokemonService services.PokemonService
}

func NewPokemonHandler(pokemonService services.PokemonService) PokemonHandler {
	return PokemonHandler{pokemonService: pokemonService}
}

func (h PokemonHandler) GetAllPokemons(c *gin.Context) {
	search := c.Query("search")
	var pokemons models.Resource
	var err error

	if search != "" {
		pokemons, err = h.pokemonService.SearchPokemons(search)
	} else {
		pokemons, err = h.pokemonService.GetAllPokemons()
	}

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"pokemons": pokemons,
	})

}

func (h PokemonHandler) PokemonDetail(c *gin.Context) {
	pokemonId := c.Param("id")

	pokemon, err := h.pokemonService.PokemonDetail(pokemonId)

	if shared.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"pokemon": pokemon,
	})

}
