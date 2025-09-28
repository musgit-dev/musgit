package models

import "slices"

type User struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Pieces    []*Piece    `json:"pieces"`
	Practices []*Practice `json:"practices"`
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func (u *User) AssignPiece(piece *Piece) {
	if !slices.Contains(u.Pieces, piece) {
		u.Pieces = append(u.Pieces, piece)
	}
}

func (u *User) PracticePiece(practice *Practice, lesson *Lesson) {
	if !slices.Contains(u.Practices, practice) {
		u.Practices = append(u.Practices, practice)
	}
}
