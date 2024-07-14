package api

import (
	"log"
	"math/rand"
)

// CreateInitialTeams creates the initial set of teams if they don't already exist
func CreateInitialTeams() {
	log.Println("Creating initial teams...")

	var count int64
	DB.Model(&Team{}).Count(&count)

	if count > 0 {
		log.Println("Teams already exist in the database. Skipping creation.")
		return
	}

	teams := []Team{
		{Name: "Chelsea", Power: GenerateRandomPower(0, 100)},
		{Name: "Arsenal", Power: GenerateRandomPower(0, 100)},
		{Name: "Manchester City", Power: GenerateRandomPower(0, 100)},
		{Name: "Liverpool", Power: GenerateRandomPower(0, 100)},
	}

	for _, team := range teams {
		DB.Create(&team)
	}

	log.Println("Initial teams created successfully")
}

// GenerateRandomPower generates a random power value between min and max
func GenerateRandomPower(min, max int) int {
	return rand.Intn(max-min+1) + min
}
