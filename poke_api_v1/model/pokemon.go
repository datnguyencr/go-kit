package model

type Pokemon struct {
	Abilities        []string `json:"abilities"`
	AgainstBug       float64  `json:"against_bug"`
	AgainstDark      float64  `json:"against_dark"`
	AgainstDragon    float64  `json:"against_dragon"`
	AgainstElectric  float64  `json:"against_electric"`
	AgainstFairy     float64  `json:"against_fairy"`
	AgainstFight     float64  `json:"against_fight"`
	AgainstFire      float64  `json:"against_fire"`
	AgainstFlying    float64  `json:"against_flying"`
	AgainstGhost     float64  `json:"against_ghost"`
	AgainstGrass     float64  `json:"against_grass"`
	AgainstGround    float64  `json:"against_ground"`
	AgainstIce       float64  `json:"against_ice"`
	AgainstNormal    float64  `json:"against_normal"`
	AgainstPoison    float64  `json:"against_poison"`
	AgainstPsychic   float64  `json:"against_psychic"`
	AgainstRock      float64  `json:"against_rock"`
	AgainstSteel     float64  `json:"against_steel"`
	AgainstWater     float64  `json:"against_water"`
	Attack           int      `json:"attack"`
	BaseEggSteps     int      `json:"base_egg_steps"`
	BaseHappiness    int      `json:"base_happiness"`
	BaseTotal        int      `json:"base_total"`
	CaptureRate      int      `json:"capture_rate"`
	Classification   string   `json:"classification"`
	Defense          int      `json:"defense"`
	ExperienceGrowth int      `json:"experience_growth"`
	Height           float64  `json:"height"`
	HP               int      `json:"hp"`
	JapaneseName     string   `json:"japanese_name"`
	Name             string   `json:"name"`
	PercentageMale   float64  `json:"percentage_male"`
	PokedexNumber    int      `json:"pokedex_number"`
	SpAttack         int      `json:"sp_attack"`
	SpDefense        int      `json:"sp_defense"`
	Speed            int      `json:"speed"`
	Type1            string   `json:"type1"`
	Type2            string   `json:"type2"`
	Weight           float64  `json:"weight"`
	Generation       int      `json:"generation"`
	IsLegendary      int      `json:"is_legendary"`
}
