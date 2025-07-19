package domain

import (
	"testing"
)

func TestPiece(t *testing.T) {

	// New practice
	piece := NewPiece("test", "test composer", Easy)

	if len(piece.Practices) != 0 {
		t.Fatal("Unexpected practices")
	}

	// Start practice
	practice, err := piece.StartPractice()
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

	_, err = piece.StartPractice()
	if err == nil {
		t.Fatal("Not allowed to create another practice")
	}

	// Complete practice
	practice, _ = piece.StopPractice(Normal)
	if practice.EndDate.IsZero() {
		t.Fatal("Practice should have ended")
	}
	if practice.Progress != Normal {
		t.Fatal("Practice evaluation mismatch")
	}

	_, err = piece.StopPractice(Normal)
	if err == nil {
		t.Fatal("Not allowed to complete not started practice")
	}

}
