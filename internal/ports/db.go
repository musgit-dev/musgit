package ports

import "musgit/internal/application/domain"

type DBPort interface {
	AddPiece(piece *domain.Piece) (*domain.Piece, error)
	GetPiece(id int64) (domain.Piece, error)
	// GetPieces() ([]domain.Piece, error)
}
