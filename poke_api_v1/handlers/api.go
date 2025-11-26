package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"poke_api/model"

	"github.com/gin-gonic/gin"
)

type Pokemon = model.Pokemon

func APIHandler(c *gin.Context) {
	entries, err := LoadPokemonCSV("./db/pokemon.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)
}

func LoadPokemonCSV(path string) ([]Pokemon, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("csv has no data")
	}

	var pokemons []Pokemon
	for i, row := range rows[1:] {
		if len(row) < 36 {
			fmt.Printf("row %d skipped: not enough columns\n", i+1)
			continue
		}

		p := Pokemon{
			Abilities:        parseAbilities(row[0]),
			AgainstBug:       safeAtof(row[1]),
			AgainstDark:      safeAtof(row[2]),
			AgainstDragon:    safeAtof(row[3]),
			AgainstElectric:  safeAtof(row[4]),
			AgainstFairy:     safeAtof(row[5]),
			AgainstFight:     safeAtof(row[6]),
			AgainstFire:      safeAtof(row[7]),
			AgainstFlying:    safeAtof(row[8]),
			AgainstGhost:     safeAtof(row[9]),
			AgainstGrass:     safeAtof(row[10]),
			AgainstGround:    safeAtof(row[11]),
			AgainstIce:       safeAtof(row[12]),
			AgainstNormal:    safeAtof(row[13]),
			AgainstPoison:    safeAtof(row[14]),
			AgainstPsychic:   safeAtof(row[15]),
			AgainstRock:      safeAtof(row[16]),
			AgainstSteel:     safeAtof(row[17]),
			AgainstWater:     safeAtof(row[18]),
			Attack:           safeAtoi(row[19]),
			BaseEggSteps:     safeAtoi(row[20]),
			BaseHappiness:    safeAtoi(row[21]),
			BaseTotal:        safeAtoi(row[22]),
			CaptureRate:      safeAtoi(row[23]),
			Classification:   row[24],
			Defense:          safeAtoi(row[25]),
			ExperienceGrowth: safeAtoi(row[26]),
			Height:           safeAtof(row[27]),
			HP:               safeAtoi(row[28]),
			JapaneseName:     row[29],
			Name:             row[30],
			PercentageMale:   safeAtof(row[31]),
			PokedexNumber:    safeAtoi(row[32]),
			SpAttack:         safeAtoi(row[33]),
			SpDefense:        safeAtoi(row[34]),
			Speed:            safeAtoi(row[35]),
			Type1:            row[36],
			Type2:            row[37],
			Weight:           safeAtof(row[38]),
			Generation:       safeAtoi(row[39]),
			IsLegendary:      safeAtoi(row[40]),
		}

		pokemons = append(pokemons, p)
	}

	return pokemons, nil
}
func safeIndex(row []string, idx int) string {
	if idx >= len(row) {
		return ""
	}
	return row[idx]
}
func parseAbilities(s string) []string {
	s = strings.Trim(s, "[]\"'") // remove brackets and quotes
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.Trim(parts[i], " '\"")
	}
	return parts
}

func safeAtoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0
	}
	return i
}

func safeAtof(s string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0
	}
	return f
}
