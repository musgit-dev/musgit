package db

import (
	"fmt"
	"musgit/internal/application/domain"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Piece struct {
	gorm.Model
	Name            string
	Composer        domain.Composer
	PieceComplexity domain.PieceComplexity
	State           domain.PieceState
}

type Practice struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
	Progress  domain.PracticeProgressEvalutation
	PieceId   uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("Db connection error: %v", openErr)
	}
	if err := db.AutoMigrate(&Piece{}, &Practice{}); err != nil {
		return nil, fmt.Errorf("Db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}
