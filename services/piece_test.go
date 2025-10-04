package services

import (
	"testing"

	"github.com/musgit-dev/musgit/models"
)

func TestAddPiece(t *testing.T) {
	piece, err := pieceService.Add(
		"piece",
		"test_composer",
		models.PieceComplexityEasy,
	)
	if err != nil {
		t.Fatal(err)
	}

	if id := piece.ID; id != 2 {
		t.Fatalf("Unknown piece id, should be 2, got %d", id)
	}

	newPiece, err := pieceService.Get(2)
	if err != nil {
		t.Fatal(err)
	}

	if name := newPiece.Name; name != "piece" {
		t.Fatalf("Unknown piece name, expected 'test', got %s", name)
	}
}

func TestGetAllPieces(t *testing.T) {

	pieces := pieceService.GetAll()

	if nPieces := len(pieces); nPieces != 2 {
		t.Fatalf("Unexpected number of pieces, got %d, expected 2", nPieces)
	}
}
