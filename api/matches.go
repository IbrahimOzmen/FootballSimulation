package api

import (
	"log"
)

func GetMatchesByWeek(week int) []Match {
	var matches []Match
	result := DB.Preload("HomeTeam").Preload("AwayTeam").Where("week = ?", week).Find(&matches)
	if result.Error != nil {
		log.Printf("Error retrieving matches for week %d: %v", week, result.Error)
		return nil
	}
	return matches
}
