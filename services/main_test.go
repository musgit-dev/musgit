package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/musgit-dev/musgit/internal/adapters/db"
	"github.com/musgit-dev/musgit/models"
)

var pieceService *PieceService

func TestMain(m *testing.M) {
	dbPort, err := db.NewAdapter(":memory:")
	if err != nil {
		os.Exit(1)
	}
	lessonService := NewLessonService(dbPort)
	pieceService = NewPieceService(dbPort)
	practiceService = NewPracticeService(dbPort)
	userService = NewUserService(dbPort)
	piece, err := pieceService.Add(
		"test_piece",
		"test_composer",
		models.PieceComplexityEasy,
	)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Added piece %d\n", piece.ID)
	lesson, err := lessonService.Start()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Started lesson %d\n", lesson.ID)

	exitVal := m.Run()
	os.Exit(exitVal)
}
