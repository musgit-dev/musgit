package services

import (
	"testing"
)

var userService *UserService

func TestAdd(t *testing.T) {
	user, err := userService.Add("test_user")
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "test_user" {
		t.Fatalf("Unexpected user name, got %s, expected test user", user.Name)
	}
}

func TestAssignPiece(t *testing.T) {
	user, err := userService.Add("test_user")
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
