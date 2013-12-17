package lib

import (
	"bytes"
	"strings"
)

// The representation of a TTT board
type Board struct {
	// A slice of the board's cells
	Cells []*Cell
	// Width of the board
	Width int
	// Height of the board
	Height int
	// Number of pieces in a row required to win
	MatchesToWin int
	// The winner (if any)
	Winner string
}

// Generate a string representation of the board
func (b *Board) String() string {
	var buffer bytes.Buffer
	var rowDiv string = strings.Repeat("-", (b.Width*2)+1)

	buffer.WriteString(rowDiv)
	for i := range b.Cells {

		if i%b.Width == 0 {
			buffer.WriteString("\n")
			buffer.WriteString("|")
		}

		if b.Cells[i].IsOccupied() {
			buffer.WriteString(b.Cells[i].String())
		} else {
			buffer.WriteString(" ")
		}

		buffer.WriteString("|")
	}
	buffer.WriteString("\n")
	buffer.WriteString(rowDiv)

	return buffer.String()
}

// Place a piece in the indexed cell
func (b *Board) PlacePiece(piece string, index int) {
	b.Cells[index].SetValue(piece)
}

// Create a new board with the given dimensions and settings
func NewBoard(height, width, matches_to_win int) *Board {
	board := new(Board)
	board.Cells = make([]*Cell, height*width)
	for i := range board.Cells {
		board.Cells[i] = new(Cell)
	}
	board.Height = height
	board.Width = width
	board.MatchesToWin = matches_to_win

	return board
}

// Checks the win conditions for the board
func (b *Board) IsGameOver() bool {
	piece := "X"

	// We don't support diagonals
	if b.hasHorizontalWin(piece) || b.hasVerticalWin(piece) {
		b.Winner = piece
		return true
	}

	return b.isFull()
}

// Checks for a horizontal win condition for the provided piece
func (b *Board) hasHorizontalWin(piece string) bool {
	matches := 0
	for i := 0; i < b.Height; i += 1 {
		for j := 0; j < b.Width; j += 1 {
			// K is the current cell we're checking
			k := j + (i * b.Width)
			if b.Cells[k].IsOccupiedBy(piece) {
				matches += 1
			} else {
				matches = 0
			}

			if matches >= b.MatchesToWin {
				b.Winner = piece
				return true
			}
		}
		matches = 0
	}

	return false
}

// Checks for a vertical win condition for the provided piece
func (b *Board) hasVerticalWin(piece string) bool {
	matches := 0
	for j := 0; j < b.Width; j += 1 {
		for i := 0; i < b.Height; i += 1 {
			k := j + (i * b.Width)
			if b.Cells[k].IsOccupiedBy(piece) {
				matches += 1
			} else {
				matches = 0
			}

			if matches >= b.MatchesToWin {
				b.Winner = piece
				return true
			}
		}
		matches = 0
	}

	return false
}

// Checks whether the board is completely occupied (a tie)
func (b *Board) isFull() bool {
	for i := range b.Cells {
		if !b.Cells[i].IsOccupied() {
			return false
		}
	}
	return true
}
