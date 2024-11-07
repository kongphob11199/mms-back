package postgres

import (
	"fmt"
	"log"
	"mms/internal/models"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed error connect to database")
	}

	err = db.AutoMigrate(&models.ModelUser{})

	if err != nil {
		defer log.Fatal("Database connection failed")
		log.Fatal("error + ", zap.Error(err))

		return nil, fmt.Errorf("faield connected to database")
	}

	return db, nil
}
