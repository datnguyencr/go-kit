package cache

import (
	"poke_api/model"
	"sync"
)

var (
	CachedHTML   string
	CacheMutex   sync.RWMutex
	PokemonCache []model.Pokemon
)
