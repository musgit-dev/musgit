package musgit

import (
	"log"

	"github.com/musgit-dev/musgit/internal/adapters/db"
	"github.com/musgit-dev/musgit/internal/ports"
	"github.com/musgit-dev/musgit/services"
)

type Musgit struct {
	db       ports.DBPort
	Practice services.PracticeService
	Lesson   services.LessonService
	Piece    services.PieceService
	User     services.UserService
}

func New(dbUri string) *Musgit {
	dbAdapter, err := db.NewAdapter(dbUri)
	if err != nil {
		log.Fatalf("Failed to init database, err: %v", err)
	}
	practiceService := services.NewPracticeService(dbAdapter)
	lessonService := services.NewLessonService(dbAdapter)
	pieceService := services.NewPieceService(dbAdapter)
	userService := services.NewUserService(dbAdapter)
	return &Musgit{
		db:       dbAdapter,
		Practice: *practiceService,
		Lesson:   *lessonService,
		Piece:    *pieceService,
		User:     *userService,
	}
}
