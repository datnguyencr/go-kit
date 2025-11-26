package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context) {
	filename := filepath.Base(c.Param("filename"))
	full := filepath.Join("./uploads", filename)

	if _, err := os.Stat(full); err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	c.FileAttachment(full, filename)
}
