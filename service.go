package musgit

import (
	"log"
	"musgit/internal/adapters/db"
	"musgit/internal/application/domain"
	"musgit/internal/ports"
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

func (m *MusgitService) StartLesson() (*domain.Lesson, error) {
	lesson := domain.NewLesson()
	lesson, err := m.db.AddLesson(lesson)
	if err != nil {
		return &domain.Lesson{}, err
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

func (m *MusgitService) GetLessons() []domain.Lesson {
	return m.db.GetLessons()
}

func (m *MusgitService) AddPiece(
	name, composer string,
	complexity domain.PieceComplexity,
) (domain.Piece, error) {
	piece := domain.NewPiece(name, composer, complexity)
	piece, err := m.db.AddPiece(piece)
	if err != nil {
		return domain.Piece{}, err
	}
	return *piece, nil
}

func (m *MusgitService) PracticePiece(
	pieceId int64,
	lessonId int64,
) (domain.Practice, error) {
	piece, err := m.db.GetPiece(pieceId)
	if err != nil {
		return domain.Practice{}, err
	}
	practice, err := piece.StartPractice()
	if err != nil {
		return domain.Practice{}, err
	}
	practice, err = m.db.AddPractice(practice, pieceId, lessonId)
	return *practice, err
}
