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

func (m *MusgitService) StartPractice(
	piece domain.Piece,
) (domain.Piece, error) {
	err := m.db.SavePiece(&piece)
	if err != nil {
		return domain.Piece{}, err
	}
	return piece, nil
}
