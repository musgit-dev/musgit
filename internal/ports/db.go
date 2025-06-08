package ports

import "musgit/internal/application/domain"

type DBPort interface {
	AddPiece(piece *domain.Piece) (*domain.Piece, error)
	GetPiece(id int64) (domain.Piece, error)
	AddPractice(
		practice *domain.Practice,
		pieceId int64,
	) (*domain.Practice, error)
	// GetPieces() ([]domain.Piece, error)
	UpdatePractice(practice *domain.Practice) error
}
