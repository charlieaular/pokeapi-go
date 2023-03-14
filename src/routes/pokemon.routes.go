package routes

import (
	"pokeapi-go/src/handlers"
	"pokeapi-go/src/repositories"
	"pokeapi-go/src/services"

	"github.com/gin-gonic/gin"
)

func RegisterPokemonRoutes(router *gin.Engine) {
	pokemonRepo := repositories.NewAddressRepo()
	pokemonService := services.NewPokemonService(pokemonRepo)
	pokemonHandler := handlers.NewPokemonHandler(pokemonService)

	pokemonRoutes := router.Group("pokemon")
	{
		pokemonRoutes.GET("", pokemonHandler.GetAllPokemons)
		pokemonRoutes.GET("/:id", pokemonHandler.PokemonDetail)
	}
}
