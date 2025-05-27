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
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	Complexity PieceComplexity `json:"complexity"`
	State      PieceState      `json:"state"`
	Practices  []Practice      `json:"practices"`
}

func NewPiece(name string, complexity PieceComplexity) *Piece {
	return &Piece{
		Name:       name,
		Complexity: complexity,
	}
}
