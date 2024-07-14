package api

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=football_simulation sslmode=disable password=1234")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Veritabanı tablolarını oluştur
	DB.AutoMigrate(&Team{}, &Match{}, &Settings{}, &ChampionshipPrediction{})
	fmt.Println("Database connection established")
}
