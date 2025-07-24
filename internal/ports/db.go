package ports

import "github.com/musgit-dev/musgit/models"

type DBPort interface {
	// Piece
	AddPiece(piece *models.Piece) (*models.Piece, error)
	GetPiece(id int64) (models.Piece, error)
	UpdatePiece(p *models.Piece) error
	GetPieces() []models.Piece
	// Lesson
	AddLesson(l *models.Lesson) (*models.Lesson, error)
	GetLastLesson() (models.Lesson, error)
	GetLesson(id int64) (models.Lesson, error)
	GetLessons() []models.Lesson
	UpdateLesson(l *models.Lesson) error
	// Practice
	AddPractice(
		practice *models.Practice,
		pieceId, lessonId int64,
	) (*models.Practice, error)
	UpdatePractice(practice *models.Practice) error
	GetPractice(id int64) (models.Practice, error)
	GetPractices() []models.Practice
}
