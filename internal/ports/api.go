package ports

import (
	"context"
	"musgit/internal/application/domain"
)

type APIPort interface {
	StartPractice(ctx context.Context, piece domain.Piece)
}
