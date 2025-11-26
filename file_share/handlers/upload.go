package handlers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	os.MkdirAll("./uploads", os.ModePerm)

	form, err := c.MultipartForm()
	if err != nil {
		c.String(400, "Invalid form: %s", err)
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		c.String(400, "No files uploaded")
		return
	}

	for _, file := range files {
		safe := filepath.Base(file.Filename)
		dst := filepath.Join("./uploads", safe)
		c.SaveUploadedFile(file, dst)
	}

	c.Redirect(302, "/")
}
