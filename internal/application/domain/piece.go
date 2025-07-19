package domain

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

func (p *Piece) currentPractice() (*Practice, error) {
	if len(p.Practices) == 0 {
		return NewPractice(), errors.New("No practices")
	}
	pr := p.Practices[len(p.Practices)-1]
	return pr, nil
}

func (p *Piece) StartPractice() (*Practice, error) {
	pr, err := p.currentPractice()
	if err == nil && pr.Active() {
		return pr, errors.New("You have an active practices.")
	}
	p.Practices = append(p.Practices, pr)
	return pr, nil
}

func (p *Piece) StopPractice(
	evaluation PracticeProgressEvalutation,
) (*Practice, error) {
	pr, err := p.currentPractice()
	if err != nil {
		return pr, err
	}
	if pr.Completed() {
		return pr, errors.New("Practice has already been completed.")
	}
	err = pr.Complete(evaluation)
	return pr, err
}
