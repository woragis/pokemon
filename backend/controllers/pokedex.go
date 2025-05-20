package controllers

type PokedexEntryInput struct {
	Caught          bool   `json:"caught"`
	Shiny           bool   `json:"shiny"`
	LivingDex       bool   `json:"living_dex"`
	ShinyLivingDex  bool   `json:"shiny_living_dex"`
	Notes           string `json:"notes"`
}

// // GET /me/pokedex/:game_id
// func GetPokedexByGame(c *fiber.Ctx) error

// // POST /me/pokedex/:game_id/:pokemon_id
// func UpsertPokedexEntry(c *fiber.Ctx) error

// // PUT /me/pokedex/entry/:entry_id
// func UpdatePokedexEntry(c *fiber.Ctx) error

// // DELETE /me/pokedex/entry/:entry_id
// func DeletePokedexEntry(c *fiber.Ctx) error
