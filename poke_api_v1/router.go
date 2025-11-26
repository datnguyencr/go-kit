package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"poke_api/handlers"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed assets/*
var assetsFS embed.FS

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 10 << 20 // 10MB

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFS(templatesFS, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	assetsSub, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/assets", http.FS(assetsSub))

	// Routes
	r.GET("/", handlers.IndexHandler)
	r.GET("/api/pokemon", handlers.APIHandler)

	return r
}
