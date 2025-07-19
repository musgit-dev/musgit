package ports

import (
	"musgit/internal/application/domain"
)

type APIPort interface {
	AddPiece(name, composer string, complexity domain.PieceComplexity)
	PracticePiece(piece domain.Piece)
}
