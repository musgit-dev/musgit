package services

import (
	"github.com/musgit-dev/musgit/internal/ports"
	"github.com/musgit-dev/musgit/models"
)

type LessonService struct {
	db ports.DBPort
}

func NewLessonService(db ports.DBPort) *LessonService {
	return &LessonService{db: db}
}

func (s *LessonService) Start() (*models.Lesson, error) {
	lesson := models.NewLesson()
	lesson, err := s.db.AddLesson(lesson)
	if err != nil {
		return &models.Lesson{}, err
	}
	return lesson, nil
}

func (s *LessonService) PauseCurrent() error {
	lesson, err := s.db.GetLastLesson()
	if err != nil {
		return err
	}
	lesson.Pause()
	err = s.db.UpdateLesson(&lesson)
	if err != nil {
		return err
	}
	return nil
}

func (s *LessonService) ResumeCurrent() error {
	lesson, err := s.db.GetLastLesson()
	if err != nil {
		return err
	}
	lesson.Resume()
	err = s.db.UpdateLesson(&lesson)
	if err != nil {
		return err
	}
	return nil
}

func (s *LessonService) StopCurrent() error {
	lesson, err := s.db.GetLastLesson()
	if err != nil {
		return err
	}
	lesson.Finish()
	err = s.db.UpdateLesson(&lesson)
	if err != nil {
		return err
	}
	return nil
}

func (s *LessonService) Get(id int64) (models.Lesson, error) {
	lesson, err := s.db.GetLesson(id)
	if err != nil {
		return models.Lesson{}, err
	}
	return lesson, nil
}
func (s *LessonService) GetAll() []models.Lesson {
	return s.db.GetLessons()
}
