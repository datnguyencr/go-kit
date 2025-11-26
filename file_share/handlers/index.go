package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	entries := listFiles("./uploads")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Content": entries,
	})
}
