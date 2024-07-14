package api

import (
	"log"
)

// ResetTeamStatsAndMatches resets the statistics of all teams and clears all matches
func ResetTeamStatsAndMatches() {
	log.Println("Starting to reset team stats and clear matches...")

	var teams []Team
	result := DB.Find(&teams)
	if result.Error != nil {
		log.Println("Error fetching teams:", result.Error)
		return
	}

	log.Printf("Found %d teams to reset", len(teams))

	for i := range teams {
		log.Printf("Resetting stats for team: %s", teams[i].Name)
		teams[i].Points = 0
		teams[i].Wins = 0
		teams[i].Draws = 0
		teams[i].Losses = 0
		teams[i].GoalsFor = 0
		teams[i].GoalsAgainst = 0
		teams[i].ChampionshipProbability = 0
		teams[i].GoalDifference = 0
		teams[i].Played = 0
	}

	// Güncellenmiş takımları ayrı ayrı kaydet
	for _, team := range teams {
		result = DB.Save(&team)
		if result.Error != nil {
			log.Println("Error saving team stats:", result.Error)
		} else {
			log.Printf("Successfully reset stats for team: %s", team.Name)
		}
	}

	// Maçları temizle
	result = DB.Exec("DELETE FROM matches")
	if result.Error != nil {
		log.Println("Error clearing matches:", result.Error)
	} else {
		log.Println("All matches cleared successfully")
	}

	log.Println("All team stats reset and matches cleared successfully")
}

func ResetChampionshipPredictions() {
	err := DB.Exec("DELETE FROM championship_predictions").Error
	if err != nil {
		log.Fatalf("Failed to reset championship_predictions table: %v", err)
	} else {
		log.Println("Championship_predictions table reset successfully.")
	}
}

func ResetGoalDifference() {
	err := DB.Exec("UPDATE teams SET goal_difference = 0").Error
	if err != nil {
		log.Fatalf("Failed to reset goal_difference field: %v", err)
	} else {
		log.Println("Goal_difference field reset successfully.")
	}
}
