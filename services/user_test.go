package services

import (
	"testing"
)

var userService *UserService

func TestUserAdd(t *testing.T) {
	user, err := userService.Add("test_user")
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "test_user" {
		t.Fatalf("Unexpected user name, got %s, expected test user", user.Name)
	}
	if userId := user.ID; userId != 1 {
		t.Fatalf("Unexpected user id, got %d, expected test user", userId)
	}

	newUser, err := userService.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	if newUser.Name != user.Name {
		t.Fatal("unexpected user")
	}
	users, err := userService.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if nUsers := len(users); nUsers != 1 {
		t.Fatalf("Unexpected number of users, expected 1, got %d", nUsers)
	}

}

func TestUserAssignPiece(t *testing.T) {
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

func TestUserStartPractice(t *testing.T) {
	user, _ := userService.Add("test_user")
	if nPractices := len(user.Practices); nPractices != 0 {
		t.Fatalf(
			"Unexpected number of practices, got %d, expected 0",
			nPractices,
		)
	}
	practice, err := userService.StartPractice(user.ID, 1, 1)

	if err != nil {
		t.Fatal(err)
	}

	if !practice.Active() {
		t.Fatal("Practice is not active")
	}

}
