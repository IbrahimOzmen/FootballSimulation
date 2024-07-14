package api

import (
	"log"
	"math"
	"math/rand"

	"github.com/jinzhu/gorm"
)

type Simulator struct {
	DB *gorm.DB
}

func NewSimulator(db *gorm.DB) *Simulator {
	return &Simulator{DB: db}
}

// SimulateMatches simulates the matches for the given week
func (s *Simulator) SimulateMatches(week int) {
	log.Printf("Simulating matches for week %d", week)

	var teams []Team
	s.DB.Find(&teams)

	if len(teams) < 4 {
		log.Println("Not enough teams to simulate matches")
		return
	}

	var matches []Match
	switch week {
	case 1:
		matches = []Match{
			{HomeTeamID: teams[0].ID, AwayTeamID: teams[1].ID, Week: week},
			{HomeTeamID: teams[2].ID, AwayTeamID: teams[3].ID, Week: week},
		}
	case 2:
		matches = []Match{
			{HomeTeamID: teams[0].ID, AwayTeamID: teams[2].ID, Week: week},
			{HomeTeamID: teams[1].ID, AwayTeamID: teams[3].ID, Week: week},
		}
	case 3:
		matches = []Match{
			{HomeTeamID: teams[0].ID, AwayTeamID: teams[3].ID, Week: week},
			{HomeTeamID: teams[1].ID, AwayTeamID: teams[2].ID, Week: week},
		}
	case 4:
		matches = []Match{
			{HomeTeamID: teams[1].ID, AwayTeamID: teams[0].ID, Week: week},
			{HomeTeamID: teams[3].ID, AwayTeamID: teams[2].ID, Week: week},
		}
	case 5:
		matches = []Match{
			{HomeTeamID: teams[2].ID, AwayTeamID: teams[0].ID, Week: week},
			{HomeTeamID: teams[3].ID, AwayTeamID: teams[1].ID, Week: week},
		}
	case 6:
		matches = []Match{
			{HomeTeamID: teams[3].ID, AwayTeamID: teams[0].ID, Week: week},
			{HomeTeamID: teams[2].ID, AwayTeamID: teams[1].ID, Week: week},
		}
	default:
		log.Println("Invalid week")
		return
	}

	for i := range matches {
		// Simulate the match
		s.SimulateMatch(&matches[i])
		s.DB.Create(&matches[i])
	}

	log.Printf("Successfully simulated matches for week %d", week)
}

// SimulateMatch simulates a match between two teams and updates the teams' statistics
func (s *Simulator) SimulateMatch(match *Match) {
	homeAdvantage := 1.1

	var homeTeam, awayTeam Team
	s.DB.First(&homeTeam, match.HomeTeamID)
	s.DB.First(&awayTeam, match.AwayTeamID)

	// Calculate the expected number of goals for each team
	homeGoalsExpected := float64(homeTeam.Power) * homeAdvantage / float64(homeTeam.Power+awayTeam.Power) * 1.5
	awayGoalsExpected := float64(awayTeam.Power) / float64(homeTeam.Power+awayTeam.Power) * 1.5

	// Simulate goals for each team using Poisson distribution
	homeGoals := poisson(homeGoalsExpected)
	awayGoals := poisson(awayGoalsExpected)

	match.HomeGoals = homeGoals
	match.AwayGoals = awayGoals

	// Update points and goals
	homeTeam.GoalsFor += homeGoals
	homeTeam.GoalsAgainst += awayGoals
	awayTeam.GoalsFor += awayGoals
	awayTeam.GoalsAgainst += homeGoals

	// Update played matches count
	homeTeam.Played++
	awayTeam.Played++

	if homeGoals > awayGoals {
		homeTeam.Points += 3
		homeTeam.Wins += 1
		awayTeam.Losses += 1
	} else if awayGoals > homeGoals {
		awayTeam.Points += 3
		awayTeam.Wins += 1
		homeTeam.Losses += 1
	} else {
		homeTeam.Points += 1
		awayTeam.Points += 1
		homeTeam.Draws += 1
		awayTeam.Draws += 1
	}

	// Save the updated teams to the database
	s.DB.Save(&homeTeam)
	s.DB.Save(&awayTeam)
}

// Poisson distribution function
func poisson(lam float64) int {
	L := math.Exp(-lam)
	k := 0
	p := 1.0
	for p > L {
		k++
		p *= rand.Float64()
	}
	return k - 1
}

// DisplayLeagueTable sorts the teams by points and goal difference
// DisplayLeagueTable updates the league table and saves it to the database
func (s *Simulator) DisplayLeagueTable() {
	var teams []Team
	s.DB.Find(&teams)

	for i := range teams {
		teams[i].GoalDifference = teams[i].GoalsFor - teams[i].GoalsAgainst
		s.DB.Save(&teams[i])
	}
}
