package api

import (
	"log"
)

// SaveChampionshipPredictions saves the championship predictions to the database
func SaveChampionshipPredictions(predictions map[string]float64, week int) {
	for team, probability := range predictions {
		prediction := ChampionshipPrediction{
			TeamName:    team,
			Week:        week,
			Probability: probability,
		}
		result := DB.Create(&prediction)
		if result.Error != nil {
			log.Printf("Failed to save prediction for team %s: %v", team, result.Error)
		} else {
			log.Printf("Saved prediction for team %s: %.2f%%", team, probability)
		}
	}
}

func GetPredictionsByWeek(week int) []ChampionshipPrediction {
	var predictions []ChampionshipPrediction
	result := DB.Where("week = ?", week).Find(&predictions)
	if result.Error != nil {
		log.Printf("Error retrieving predictions for week %d: %v", week, result.Error)
		return nil
	}
	return predictions
}
