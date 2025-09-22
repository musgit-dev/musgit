package models

import (
	"errors"
)

type PieceComplexity int
type PieceState int

const (
	PieceComplexityUnknown PieceComplexity = iota
	PieceComplexityEasy
	PieceComplexityMid
	PieceComplexityHard
)

const (
	Learning PieceState = iota
	Ready
)

type Piece struct {
	Composer   Composer
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	Complexity PieceComplexity `json:"complexity"`
	State      PieceState      `json:"state"`
	Practices  []*Practice     `json:"practices"`
}

func NewPiece(name, composer string, complexity PieceComplexity) *Piece {
	return &Piece{
		Composer:   Composer{Name: composer},
		Name:       name,
		Complexity: complexity,
	}
}

func (p *Piece) isCurrentlyPracticed() bool {
	if len(p.Practices) > 0 && p.Practices[len(p.Practices)-1].Active() {
		return true
	}
	return false
}

func (p *Piece) StartPractice(lessonId int64) (*Practice, error) {
	if p.isCurrentlyPracticed() {
		return nil, errors.New("You have an active practice.")
	}
	pr := NewPractice(p.ID, lessonId)
	p.Practices = append(p.Practices, pr)
	return pr, nil
}

func (p *Piece) StopPractice(
	evaluation PracticeProgressEvalutation,
) (*Practice, error) {
	if !p.isCurrentlyPracticed() {
		return nil, errors.New("You don't have an active practice.")
	}
	pr := p.Practices[len(p.Practices)-1]
	err := pr.Complete(evaluation)
	return pr, err
}
