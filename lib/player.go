package lib

type player struct {
	piece string
	score int
}

func NewPlayer(piece string) *player {
	return &player{piece: piece, score: 0}
}

func (p *player) Piece() string {
	return p.piece
}

func (p *player) SetPiece(piece string) {
	p.piece = piece
}

func (p *player) Score() int {
	return p.score
}

func (p *player) IncrementScore() {
	p.score += 1
}
