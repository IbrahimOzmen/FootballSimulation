package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTeams retrieves all teams from the database
func GetTeams(c *gin.Context) {
	var teams []Team
	DB.Find(&teams)
	c.JSON(http.StatusOK, teams)
}

// GetTeam retrieves a specific team by ID from the database
func GetTeam(c *gin.Context) {
	id := c.Param("id")
	var team Team
	if err := DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}
