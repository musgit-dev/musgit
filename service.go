package musgit

import (
	"fmt"
	"log"

	"github.com/musgit-dev/musgit/internal/adapters/db"
	"github.com/musgit-dev/musgit/internal/ports"
	"github.com/musgit-dev/musgit/models"
)

type MusgitService struct {
	db ports.DBPort
}

func NewMusgitService(dbUri string) *MusgitService {
	dbAdapter, err := db.NewAdapter(dbUri)
	if err != nil {
		log.Fatalf("Failed to init database, err: %v", err)
	}
	return &MusgitService{db: dbAdapter}
}

func (m *MusgitService) StartLesson() (*models.Lesson, error) {
	lesson := models.NewLesson()
	lesson, err := m.db.AddLesson(lesson)
	if err != nil {
		return &models.Lesson{}, err
	}
	return lesson, nil
}

func (m *MusgitService) PauseCurrentLesson() error {
	lesson, err := m.db.GetLastLesson()
	if err != nil {
		return err
	}
	lesson.Pause()
	err = m.db.UpdateLesson(&lesson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MusgitService) ResumeCurrentLesson() error {
	lesson, err := m.db.GetLastLesson()
	if err != nil {
		return err
	}
	lesson.Resume()
	err = m.db.UpdateLesson(&lesson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MusgitService) StopCurrentLesson() error {
	lesson, err := m.db.GetLastLesson()
	if err != nil {
		return err
	}
	lesson.Finish()
	err = m.db.UpdateLesson(&lesson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MusgitService) GetLessons() []models.Lesson {
	return m.db.GetLessons()
}

func (m *MusgitService) GetPieces() []models.Piece {
	return m.db.GetPieces()
}

func (m *MusgitService) GetPiece(id int64) (models.Piece, error) {
	piece, err := m.db.GetPiece(id)
	if err != nil {
		return models.Piece{}, fmt.Errorf("Unknown piece id: %d", id)
	}
	return piece, nil
}

func (m *MusgitService) AddPiece(
	name, composer string,
	complexity models.PieceComplexity,
) (models.Piece, error) {
	piece := models.NewPiece(name, composer, complexity)
	piece, err := m.db.AddPiece(piece)
	if err != nil {
		return models.Piece{}, err
	}
	return *piece, nil
}

func (m *MusgitService) PracticePiece(
	pieceId int64,
	lessonId int64,
) (models.Practice, error) {
	piece, err := m.db.GetPiece(pieceId)
	if err != nil {
		return models.Practice{}, err
	}
	practice, err := piece.StartPractice()
	if err != nil {
		return models.Practice{}, err
	}
	practice, err = m.db.AddPractice(practice, pieceId, lessonId)
	return *practice, err
}
