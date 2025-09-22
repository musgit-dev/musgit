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
	fmt.Printf("Added piece %d\n", piece.ID)
	lesson, err := lessonService.Start()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Started lesson %d\n", lesson.ID)

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

func TestStop(t *testing.T) {
	service.Start(1, 0)
	practice, err := service.Stop(1)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}

	if practice.Progress != models.PracticeProgressNormal {
		t.Fatalf(
			"Unexpected progress, expect %d, got %d",
			models.PracticeProgressNormal,
			practice.Progress,
		)
	}
	// _, err = service.Stop(1)
	// if err != models.ErrCompletedPractice {
	// 	t.Fatal(err)
	// }
}

func TestStartWarmup(t *testing.T) {
	warmup, err := service.Warmup(1)
	if err != nil {
		t.Fatal("Unknown error", err)
	}

	if warmup.State != models.WarmupActive {
		t.Fatal("Warmup should be active, got", warmup.State)
	}
	// warmup, err = service.Warmup(1)
	// if err != models.ErrWarmapStarted {
	// 	t.Fatalf("Expected error %s, got %s", models.ErrWarmapStarted, err)
	// }

	warmup, err = service.StopWarmup()
	if err != nil {
		t.Fatal("Unknown error", err)
	}
	if warmup.State != models.WarmupCompleted {
		t.Fatal("Warmup should be completed, got", warmup.State)
	}

}
