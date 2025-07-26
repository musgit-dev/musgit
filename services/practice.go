package services

import (
	"github.com/musgit-dev/musgit/internal/ports"
	"github.com/musgit-dev/musgit/models"
)

type PracticeService struct {
	db ports.DBPort
}

func NewPracticeService(db ports.DBPort) *PracticeService {
	return &PracticeService{db: db}
}

func (s *PracticeService) Start(
	pieceId, lessonId int64,
) (*models.Practice, error) {
	piece, err := s.db.GetPiece(pieceId)
	if err != nil {
		return &models.Practice{}, err
	}
	practice, err := piece.StartPractice()
	if err != nil {
		return &models.Practice{}, err
	}
	practice, err = s.db.AddPractice(practice)
	return practice, err
}

func (s *PracticeService) Stop(
	pieceId int64,
) (*models.Practice, error) {
	piece, err := s.db.GetPiece(pieceId)
	if err != nil {
		return &models.Practice{}, err
	}
	practice, err := piece.StopPractice(models.PracticeProgressNormal)
	if err != nil {
		return &models.Practice{}, err
	}
	err = s.db.UpdatePractice(practice)
	return practice, err
}
