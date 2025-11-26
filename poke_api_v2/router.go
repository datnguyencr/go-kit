package main

import (
	"html/template"
	"strings"

	"poke_api/handlers"
	"sync"

	"github.com/gin-contrib/gzip"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

var (
	cachedHTML string
	cacheMutex sync.RWMutex
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.MaxMultipartMemory = 10 << 20 // 10MB
	r.Static("/assets", "./assets")
	refreshCache("db/pokemon.csv")

	go watchFileChanges("db/pokemon.csv") // adjust path if needed

	// Routes
	r.GET("/", func(c *gin.Context) {
		cacheMutex.RLock()
		html := cachedHTML
		cacheMutex.RUnlock()
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})
	r.GET("/api/pokemon", handlers.APIHandler)

	return r
}

func refreshCache(path string) {
	pokemonCache, _ := handlers.LoadPokemonCSV(path)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	if err := tmpl.Execute(&sb, map[string]any{"Content": pokemonCache}); err == nil {
		cacheMutex.Lock()
		cachedHTML = sb.String()
		cacheMutex.Unlock()
	}
}

func watchFileChanges(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	if err := watcher.Add(path); err != nil {
		panic(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				refreshCache(path)
			}
		case err := <-watcher.Errors:
			if err != nil {
				// log or ignore
			}
		}
	}

}
