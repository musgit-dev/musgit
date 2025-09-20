package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/musgit-dev/musgit/internal/adapters/db"
	"github.com/musgit-dev/musgit/models"
)

var service *PracticeService

func TestMain(m *testing.M) {
	dbPort, err := db.NewAdapter(":memory:")
	if err != nil {
		os.Exit(1)
	}
	pieceService := NewPieceService(dbPort)
	lessonService := NewLessonService(dbPort)
	service = NewPracticeService(dbPort)
	piece, err := pieceService.Add(
		"test_piece",
		"test_composer",
		models.PieceComplexityEasy,
	)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Added piece %d", piece.ID)
	lesson, err := lessonService.Start()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Started lesson %d", lesson.ID)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestStart(t *testing.T) {
	practice, err := service.Start(1, 1)
	if err != nil {
		t.Fatal("Failed to start practice", err)
	}
	if pieceId := practice.PieceID; pieceId != 1 {
		t.Fatalf("Incorrect piece id, got %d", pieceId)
	}
	if lessonId := practice.LessonID; lessonId != 1 {
		t.Fatalf("Incorrect lesson id, got %d", lessonId)
	}
	_, err = service.Start(1, 2)
	if err != models.ErrNotActiveLesson {
		t.Fatalf("Expected error %s, got %s", models.ErrNotActiveLesson, err)
	}
	_, err = service.Start(1, 0)
	if err == models.ErrNotActiveLesson {
		t.Fatalf("Unexpected error %s", err)
	}
}
