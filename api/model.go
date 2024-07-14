package api

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ChampionshipPrediction struct {
	ID          uint      `gorm:"primary_key"`
	TeamName    string    `gorm:"type:varchar(255);not null"`
	Week        int       `gorm:"not null"`
	Probability float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Team represents a football team
type Team struct {
	gorm.Model
	Name                    string  `json:"name"`
	Power                   int     `json:"power"`
	Points                  int     `json:"points"`
	Wins                    int     `json:"wins"`
	Draws                   int     `json:"draws"`
	Losses                  int     `json:"losses"`
	GoalsFor                int     `json:"goals_for"`
	GoalsAgainst            int     `json:"goals_against"`
	GoalDifference          int     `json:"goal_difference"`
	ChampionshipProbability float64 `json:"championship_probability"`
	Played                  int     `json:"played"`
}

// Match represents a football match
type Match struct {
	gorm.Model
	HomeTeamID uint `json:"home_team_id"`
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID"`
	AwayTeamID uint `json:"away_team_id"`
	AwayTeam   Team `gorm:"foreignKey:AwayTeamID"`
	HomeGoals  int  `json:"home_goals"`
	AwayGoals  int  `json:"away_goals"`
	Week       int  `json:"week"`
}
