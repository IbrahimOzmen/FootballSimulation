package api

// CalculateChampionshipProbabilities calculates the championship probabilities for each team
func (s *Simulator) CalculateChampionshipProbabilities() map[string]float64 {
	var teams []Team
	s.DB.Find(&teams)

	totalPoints := 0
	for _, team := range teams {
		totalPoints += team.Points
	}

	predictions := make(map[string]float64)
	for _, team := range teams {
		var probability float64
		if totalPoints == 0 {
			probability = 0.0
		} else {
			probability = float64(team.Points) / float64(totalPoints) * 100
		}
		predictions[team.Name] = probability

		// Update team's championship probability
		team.ChampionshipProbability = probability
		s.DB.Save(&team)
	}
	return predictions
}
