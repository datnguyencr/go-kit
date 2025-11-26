package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	cachedHTML string
	cacheMutex sync.RWMutex
)

func IndexHandler(c *gin.Context) {
	cacheMutex.RLock()
	html := cachedHTML
	cacheMutex.RUnlock()

	if html == "" {
		c.String(http.StatusInternalServerError, "No cached HTML available yet")
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))

}
