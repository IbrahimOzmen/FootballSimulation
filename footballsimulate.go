package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"FootballSimulation/api"

	"github.com/gin-gonic/gin"
)

// TemplateData struct holds the data to be passed to the HTML template
type TemplateData struct {
	Teams                   []api.Team
	Week                    int
	Matches                 []api.Match
	GroupedMatches          map[int][]api.Match
	ChampionshipPredictions map[string]float64
	PlayAll                 bool
	LeagueEnded             bool
}

// add adds two integers
func add(a, b int) int {
	return a + b
}

// subtract subtracts two integers
func subtract(a, b int) int {
	return a - b
}

// groupByWeek groups matches by their week
func groupByWeek(matches []api.Match) map[int][]api.Match {
	grouped := make(map[int][]api.Match)
	for _, match := range matches {
		week := match.Week
		grouped[week] = append(grouped[week], match)
	}
	return grouped
}

// sortByPoints sorts the teams by points and goal difference
func sortByPoints(teams []api.Team) {
	sort.SliceStable(teams, func(i, j int) bool {
		if teams[i].Points == teams[j].Points {
			return teams[i].GoalDifference > teams[j].GoalDifference
		}
		return teams[i].Points > teams[j].Points
	})
}

// HandleRoot handles the root URL and displays the league table and match results
func HandleRoot(c *gin.Context) {
	log.Println("HandleRoot called")

	weekParam := c.Query("week")
	playAllParam := c.Query("playall")

	settings := api.GetSettings()
	currentWeek := settings.CurrentWeek
	leagueEnded := settings.LeagueEnded

	week, err := strconv.Atoi(weekParam)
	if err != nil || week < 0 || week > 6 {
		week = currentWeek
	}

	simulator := api.NewSimulator(api.DB)

	playAll := playAllParam == "true"
	if playAll && !leagueEnded {
		for i := currentWeek + 1; i <= 6; i++ {
			simulator.SimulateMatches(i)
		}
		week = 6
		leagueEnded = true
		api.SaveSettings(week, leagueEnded)
	} else if week > currentWeek && !leagueEnded {
		simulator.SimulateMatches(week)
		currentWeek = week
		if currentWeek >= 6 {
			leagueEnded = true
		}
		api.SaveSettings(currentWeek, leagueEnded)
	}

	var matches []api.Match
	api.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches)

	groupedMatches := groupByWeek(matches)
	predictions := simulator.CalculateChampionshipProbabilities()

	// Güncellenmiş takımları veritabanından tekrar çek
	var teams []api.Team
	api.DB.Find(&teams)

	// Lig tablosunu güncelle ve sırala
	simulator.DisplayLeagueTable()
	sortByPoints(teams)

	// Şampiyonluk ihtimallerini kaydet
	log.Println("Saving championship predictions...")
	api.SaveChampionshipPredictions(predictions, week)

	data := TemplateData{
		Teams:                   teams,
		Week:                    week,
		Matches:                 groupedMatches[week],
		GroupedMatches:          groupedMatches,
		ChampionshipPredictions: predictions,
		PlayAll:                 playAll,
		LeagueEnded:             leagueEnded,
	}

	tmpl := template.Must(template.New("template.html").Funcs(template.FuncMap{
		"subtract":    subtract,
		"add":         add,
		"groupByWeek": groupByWeek,
	}).ParseFiles("templates/template.html"))

	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		log.Println("Error executing template:", err)
	} else {
		log.Println("Template executed successfully")
	}
}

func main() {
	log.Println("Starting application...")

	rand.Seed(time.Now().UnixNano())

	api.InitDatabase()
	api.ResetSettings()
	api.CreateInitialTeams()
	api.ResetTeamStatsAndMatches()
	api.ResetChampionshipPredictions() // Championship predictions tablosunu sıfırlıyoruz
	api.ResetGoalDifference()

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"subtract":    subtract,
		"add":         add,
		"groupByWeek": groupByWeek,
	})
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", HandleRoot)
	r.GET("/teams", api.GetTeams)
	r.GET("/teams/:id", api.GetTeam)
	r.GET("/simulate", func(c *gin.Context) {
		weekStr := c.DefaultQuery("week", "1")
		week, err := strconv.Atoi(weekStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
			return
		}
		simulator := api.NewSimulator(api.DB)
		simulator.SimulateMatches(week)
		c.JSON(http.StatusOK, gin.H{"message": "Matches simulated"})
	})

	r.GET("/weeks/:week/matches", func(c *gin.Context) {
		weekStr := c.Param("week")
		week, err := strconv.Atoi(weekStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
			return
		}
		matches := api.GetMatchesByWeek(week)
		if matches == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Matches not found"})
			return
		}
		c.JSON(http.StatusOK, matches)
	})

	r.GET("/predictions/:week", func(c *gin.Context) {
		weekStr := c.Param("week")
		week, err := strconv.Atoi(weekStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
			return
		}
		predictions := api.GetPredictionsByWeek(week)
		c.JSON(http.StatusOK, predictions)
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
