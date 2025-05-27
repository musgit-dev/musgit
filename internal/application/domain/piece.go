package domain

type PieceComplexity int
type PieceState int

const (
	Unknown PieceComplexity = iota
	Easy
	Mid
	Hard
)

const (
	Learning PieceState = iota
	Ready
)

type Piece struct {
	Name       string
	Composer   Composer
	Complexity PieceComplexity
	State      PieceState
	Practices  []Practice
}

func NewPiece(name, composer string, complexity PieceComplexity) *Piece {
	return &Piece{
		Name:       name,
		Composer:   Composer{Name: composer},
		Complexity: complexity,
	}
}
