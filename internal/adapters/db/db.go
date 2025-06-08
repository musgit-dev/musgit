package db

import (
	"errors"
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
	Composer        Composer
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

func (a *Adapter) checkPiece(name string) bool {
	var p Piece
	res := a.db.Where("name = ?", name).First(&p)
	return res.RowsAffected != 0
}

func (a *Adapter) checkComposer(name string) uint {
	var c Composer
	_ = a.db.Where("name = ?", name).First(&c)
	return c.ID
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

func (a *Adapter) AddPiece(piece *domain.Piece) (*domain.Piece, error) {

	if a.checkPiece(piece.Name) {
		return &domain.Piece{}, errors.New("Already exists")
	}

	composerId := a.checkComposer(piece.Composer.Name)

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

	if composerId != 0 {
		pieceModel.ComposerId = composerId
	} else {
		pieceModel.Composer = Composer{Name: piece.Composer.Name}
	}

	res := a.db.Create(&pieceModel)
	if res.Error == nil {
		piece.ID = int64(pieceModel.ID)
	}
	return piece, res.Error
}

func (a *Adapter) AddPractice(
	practice *domain.Practice,
	pieceId int64,
) (*domain.Practice, error) {

	practiceModel := Practice{
		StartDate: practice.StartDate,
		PieceId:   uint(pieceId),
	}

	res := a.db.Create(&practiceModel)
	if res.Error == nil {
		practice.ID = int64(practiceModel.ID)
	}
	return practice, res.Error
}

func (a *Adapter) UpdatePractice(practice *domain.Practice) error {
	res := a.db.Save(practice)
	return res.Error
}
