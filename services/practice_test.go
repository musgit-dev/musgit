package services

import (
	"testing"

	"github.com/musgit-dev/musgit/models"
)

var practiceService *PracticeService

func TestStart(t *testing.T) {
	practice, err := practiceService.Start(1, 1)
	if err != nil {
		t.Fatal("Failed to start practice", err)
	}
	if pieceId := practice.PieceID; pieceId != 1 {
		t.Fatalf("Incorrect piece id, got %d", pieceId)
	}
	if lessonId := practice.LessonID; lessonId != 1 {
		t.Fatalf("Incorrect lesson id, got %d", lessonId)
	}

	_, err = practiceService.Start(1, 2)
	if err != models.ErrNotActiveLesson {
		t.Fatalf("Expected error %s, got %s", models.ErrNotActiveLesson, err)
	}
	_, err = practiceService.Start(1, 0)
	if err == models.ErrNotActiveLesson {
		t.Fatalf("Unexpected error %s", err)
	}
}

func TestStop(t *testing.T) {
	practiceService.Start(1, 0)
	practice, err := practiceService.Stop(1)
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
	// _, err = practiceService.Stop(1)
	// if err != models.ErrCompletedPractice {
	// 	t.Fatal(err)
	// }
}

func TestStartWarmup(t *testing.T) {
	warmup, err := practiceService.Warmup(1)
	if err != nil {
		t.Fatal("Unknown error", err)
	}

	if warmup.State != models.WarmupActive {
		t.Fatal("Warmup should be active, got", warmup.State)
	}
	// warmup, err = practiceService.Warmup(1)
	// if err != models.ErrWarmapStarted {
	// 	t.Fatalf("Expected error %s, got %s", models.ErrWarmapStarted, err)
	// }

	warmup, err = practiceService.StopWarmup()
	if err != nil {
		t.Fatal("Unknown error", err)
	}
	if warmup.State != models.WarmupCompleted {
		t.Fatal("Warmup should be completed, got", warmup.State)
	}

}
