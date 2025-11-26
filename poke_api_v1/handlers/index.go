package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	entries, err := LoadPokemonCSV("./db/pokemon.csv") // capitalized if exported
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load CSV: %v", err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Content": entries,
	})
}
