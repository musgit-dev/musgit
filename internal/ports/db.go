package ports

import "musgit/internal/application/domain"

type DBPort interface {
	GetPiece(id int64) (*domain.Piece, error)
	GetPieces() ([]*domain.Piece, error)
	SavePiece(piece *domain.Piece) error
}
