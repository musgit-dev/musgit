package domain_test

import (
	"musgit/internal/application/domain"
	"testing"
)

func TestLesson(t *testing.T) {

	lesson := domain.NewLesson()

	if lesson.StartDate.IsZero() {
		t.Errorf("lesson not started")
	}

	if lesson.State != domain.LessonActive {
		t.Errorf("Expected lesson to be Active")
	}
	if !lesson.EndDate.IsZero() {
		t.Errorf("Lesson has non zero end date")
	}

	lesson.Pause()
	if lesson.State != domain.LessonPaused {
		t.Errorf("Expected lesson to be Paused")
	}

	lesson.Resume()
	if lesson.State != domain.LessonActive {
		t.Errorf("Expected lesson to be Active again")
	}

	lesson.Finish()
	if lesson.EndDate.IsZero() {
		t.Errorf("Lesson not completed")
	}
	if lesson.State != domain.LessonCompleted {
		t.Errorf("Expected lesson to be Completed")
	}

	lesson.AddNote("Good")

	if lesson.Comment != "Good" {
		t.Errorf("Comment has not been added to the lesson.")
	}

}
