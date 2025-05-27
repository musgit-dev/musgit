package api

import (
	"context"

	"musgit/internal/application/domain"
	"musgit/internal/ports"
)

type MusgitService struct {
	db ports.DBPort
}

func (m *MusgitService) StartPractice(
	ctx context.Context,
	piece domain.Piece,
) (domain.Piece, error) {
	err := m.db.SavePiece(&piece)
	if err != nil {
		return domain.Piece{}, err
	}
	return piece, nil
}
