package services

import (
	"fmt"

	"github.com/musgit-dev/musgit/internal/ports"
	"github.com/musgit-dev/musgit/models"
)

type PieceService struct {
	db ports.DBPort
}

func NewPieceService(db ports.DBPort) *PieceService {
	return &PieceService{db: db}
}

func (m *PieceService) Add(
	name, composer string,
	complexity models.PieceComplexity,
) (*models.Piece, error) {
	piece := models.NewPiece(name, composer, complexity)
	piece, err := m.db.AddPiece(piece)
	if err != nil {
		return &models.Piece{}, err
	}
	return piece, nil
}

func (s *PieceService) GetAll() []models.Piece {
	return s.db.GetPieces()
}

func (s *PieceService) Get(id int64) (models.Piece, error) {
	piece, err := s.db.GetPiece(id)
	if err != nil {
		return models.Piece{}, fmt.Errorf("Unknown piece id: %d", id)
	}
	return piece, nil
}
