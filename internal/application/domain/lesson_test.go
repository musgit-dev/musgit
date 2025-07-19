package domain

import (
	"testing"
)

func TestLesson(t *testing.T) {

	lesson := NewLesson()

	if lesson.StartDate.IsZero() {
		t.Errorf("lesson not started")
	}

	if lesson.State != LessonActive {
		t.Errorf("Expected lesson to be Active")
	}
	if !lesson.EndDate.IsZero() {
		t.Errorf("Lesson has non zero end date")
	}

	lesson.Pause()
	if lesson.State != LessonPaused {
		t.Errorf("Expected lesson to be Paused")
	}

	lesson.Resume()
	if lesson.State != LessonActive {
		t.Errorf("Expected lesson to be Active again")
	}

	lesson.Finish()
	if lesson.EndDate.IsZero() {
		t.Errorf("Lesson not completed")
	}
	if lesson.State != LessonCompleted {
		t.Errorf("Expected lesson to be Completed")
	}

	lesson.AddNote("Good")

	if lesson.Comment != "Good" {
		t.Errorf("Comment has not been added to the lesson.")
	}

}
