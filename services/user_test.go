package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/musgit-dev/musgit/internal/adapters/db"
	"github.com/musgit-dev/musgit/models"
)

var service *UserService
var pieceService *PieceService

func TestMain(m *testing.M) {
	dbPort, err := db.NewAdapter(":memory:")
	if err != nil {
		os.Exit(1)
	}
	pieceService = NewPieceService(dbPort)

	piece, err := pieceService.Add(
		"test_piece",
		"test_composer",
		models.PieceComplexityUnknown,
	)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Added piece", piece.ID)
	service = NewUserService(dbPort)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestAdd(t *testing.T) {
	user, err := service.Add("test_user")
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "test_user" {
		t.Fatalf("Unexpected user name, got %s, expected test user", user.Name)
	}
}

func TestAssignPiece(t *testing.T) {
	user, err := service.Add("test_user")
	if err != nil {
		t.Fatal(err)
	}
	piece, _ := pieceService.Get(1)
	if piece.ID != 1 {
		t.Fatal("Unknown piece")
	}
	user.AssignPiece(&piece)
	if len(user.Pieces) != 1 {
		t.Fatalf("Piece wasn't added")
	}

}
