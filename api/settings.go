package api

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Settings struct {
	gorm.Model
	CurrentWeek int
	LeagueEnded bool
}

func GetSettings() Settings {
	var settings Settings
	if err := DB.First(&settings).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Eğer ayar yoksa, varsayılan ayarları oluştur
			settings = Settings{CurrentWeek: 0, LeagueEnded: false}
			DB.Create(&settings)
		} else {
			log.Fatalf("Error retrieving settings: %v", err)
		}
	}
	return settings
}

func SaveSettings(currentWeek int, leagueEnded bool) {
	var settings Settings
	DB.First(&settings)
	settings.CurrentWeek = currentWeek
	settings.LeagueEnded = leagueEnded
	DB.Save(&settings)
}

func ResetSettings() {
	var settings Settings
	if err := DB.First(&settings).Error; err == nil {
		settings.CurrentWeek = 0
		settings.LeagueEnded = false
		DB.Save(&settings)
	} else if err == gorm.ErrRecordNotFound {
		// Eğer ayar yoksa, varsayılan ayarları oluştur
		settings = Settings{CurrentWeek: 0, LeagueEnded: false}
		DB.Create(&settings)
	} else {
		log.Fatalf("Error resetting settings: %v", err)
	}
}
