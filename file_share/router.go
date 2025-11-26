package main

import (
	"embed"
	"go_demo/handlers"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed assets/*
var assetsFS embed.FS

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 10 << 20 // 10MB

	// Use embedded templates
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	// Serve embedded static files
	r.StaticFS("/assets", http.FS(assetsFS))

	r.GET("/", handlers.IndexHandler)
	r.POST("/upload", handlers.UploadHandler)
	r.GET("/download/:filename", handlers.DownloadHandler)

	return r
}
