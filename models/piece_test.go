package models

import (
	"testing"
)

func TestPiece(t *testing.T) {

	// New practice
	piece := NewPiece("test", "test composer", PieceComplexityEasy)

	if len(piece.Practices) != 0 {
		t.Fatal("Unexpected practices")
	}

	// Start practice
	practice, err := piece.StartPractice(0)
	if err != nil {
		t.Fatal(err)
	}
	if len(piece.Practices) == 0 {
		t.Fatal("Expected practice, has none")
	}

	if !practice.Active() {
		t.Fatal("Practice should have started")
	}
	if practice.Completed() {
		t.Fatal("Practice should have not ended")
	}

	_, err = piece.StartPractice(0)
	if err == nil {
		t.Fatal("Not allowed to create another practice")
	}

	// Complete practice
	practice, _ = piece.StopPractice(PracticeProgressNormal)
	if practice.EndDate.IsZero() {
		t.Fatal("Practice should have ended")
	}
	if practice.Progress != PracticeProgressNormal {
		t.Fatal("Practice evaluation mismatch")
	}

	_, err = piece.StopPractice(PracticeProgressNormal)
	if err == nil {
		t.Fatal("Not allowed to complete not started practice")
	}

}
