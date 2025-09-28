package models

import (
	"slices"
	"testing"
)

func TestAssignPiece(t *testing.T) {
	piece := NewPiece("test piece", "test composer", PieceComplexityEasy)
	user := NewUser("test")

	user.AssignPiece(piece)
	if !slices.Contains(user.Pieces, piece) {
		t.Fatal("Piece wasn't added")
	}

}
