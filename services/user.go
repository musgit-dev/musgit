package services

import (
	"errors"

	"github.com/musgit-dev/musgit/internal/ports"
	"github.com/musgit-dev/musgit/models"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

type UserService struct {
	db ports.DBPort
}

func NewUserService(db ports.DBPort) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetAll() ([]*models.User, error) {
	users, err := s.db.GetUsers()
	if err != nil {
		return []*models.User{}, err
	}
	return users, nil
}

func (s *UserService) Get(id int64) (*models.User, error) {
	user, err := s.db.GetUser(id)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (s *UserService) Add(name string) (*models.User, error) {
	u := models.NewUser(name)
	u, err := s.db.AddUser(u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *UserService) AssignPiece(userId int64, piece models.Piece) error {
	user, err := s.Get(userId)
	if err != nil {
		return ErrUserNotFound
	}
	user.AssignPiece(&piece)
	return nil
}

func (s *UserService) StartPractice(
	userId, pieceId, lessonId int64,
) (*models.Practice, error) {
	var practice *models.Practice
	user, err := s.Get(userId)
	if err != nil {
		return practice, ErrUserNotFound
	}
	piece, err := s.db.GetPiece(pieceId)
	lesson, err := s.db.GetLesson(lessonId)
	practice, err = piece.StartPractice(lessonId)
	user.PracticePiece(practice, &lesson)
	return practice, err
}
