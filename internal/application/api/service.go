package api

import (
	"musgit/internal/application/domain"
	"musgit/internal/ports"
)

type MusgitService struct {
	db ports.DBPort
}

func NewMusgitService(db ports.DBPort) *MusgitService {
	return &MusgitService{db: db}
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

func (m *MusgitService) StartPractice(
	pieceId int64,
) (domain.Practice, error) {
	piece, err := m.db.GetPiece(pieceId)
	if err != nil {
		return domain.Practice{}, err
	}
	practice, err := piece.StartPractice()
	if err != nil {
		return domain.Practice{}, err
	}
	practice, err = m.db.AddPractice(practice, pieceId)
	return *practice, err
}

func (m *MusgitService) StopPractice(
	pieceId int64,
	evaluation domain.PracticeProgressEvalutation,
) error {
	piece, err := m.db.GetPiece(pieceId)
	if err != nil {
		return err
	}
	practice, err := piece.StopPractice(evaluation)
	if err != nil {
		return err
	}
	err = m.db.UpdatePractice(practice)
	return err
}
