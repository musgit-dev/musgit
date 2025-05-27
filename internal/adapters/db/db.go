package db

import (
	"fmt"
	"musgit/internal/application/domain"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Composer struct {
	gorm.Model
	Name string
}
type Piece struct {
	gorm.Model
	Name            string
	PieceComplexity domain.PieceComplexity
	State           domain.PieceState
	Practices       []Practice
	ComposerId      uint
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
	if err := db.AutoMigrate(&Composer{}, &Piece{}, &Practice{}); err != nil {
		return nil, fmt.Errorf("Db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) GetPiece(id int64) (domain.Piece, error) {

	var p Piece

	res := a.db.First(&p, id)
	var practices []domain.Practice

	for _, v := range p.Practices {
		practices = append(practices, domain.Practice{
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Progress:  v.Progress,
		})
	}
	piece := domain.Piece{
		ID:         int64(p.ID),
		Name:       p.Name,
		State:      p.State,
		Complexity: p.PieceComplexity,
		Practices:  practices,
	}
	return piece, res.Error
}

func (a *Adapter) SavePiece(piece *domain.Piece) error {
	var practices []Practice

	for _, v := range piece.Practices {
		practices = append(practices, Practice{
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			Progress:  v.Progress,
		})
	}
	pieceModel := Piece{
		Name:            piece.Name,
		State:           piece.State,
		PieceComplexity: piece.Complexity,
		Practices:       practices,
	}

	res := a.db.Create(&pieceModel)
	if res.Error == nil {
		piece.ID = int64(pieceModel.ID)
	}
	return res.Error

}
