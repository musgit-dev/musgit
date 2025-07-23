package ports

import "github.com/musgit-dev/musgit/internal/application/domain"

type DBPort interface {
	// Piece
	AddPiece(piece *domain.Piece) (*domain.Piece, error)
	GetPiece(id int64) (domain.Piece, error)
	UpdatePiece(p *domain.Piece) error
	GetPieces() []domain.Piece
	// Lesson
	AddLesson(l *domain.Lesson) (*domain.Lesson, error)
	GetLastLesson() (domain.Lesson, error)
	GetLesson(id int64) (domain.Lesson, error)
	GetLessons() []domain.Lesson
	UpdateLesson(l *domain.Lesson) error
	// Practice
	AddPractice(
		practice *domain.Practice,
		pieceId, lessonId int64,
	) (*domain.Practice, error)
	UpdatePractice(practice *domain.Practice) error
	GetPractice(id int64) (domain.Practice, error)
	GetPractices() []domain.Practice
}
