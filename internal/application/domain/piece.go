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
	Composer   Composer
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	Complexity PieceComplexity `json:"complexity"`
	State      PieceState      `json:"state"`
	Practices  []Practice      `json:"practices"`
}

func NewPiece(name, composer string, complexity PieceComplexity) *Piece {
	return &Piece{
		Composer:   Composer{Name: composer},
		Name:       name,
		Complexity: complexity,
	}
}

func (p *Piece) StartPractice() (*Practice, error) {
	practice := Practice{}
	p.Practices = append(p.Practices, practice)
	return &practice, nil
}
