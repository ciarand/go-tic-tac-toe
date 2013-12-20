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
	for i, cell := range b.Cells {

		if i%b.Width == 0 {
			buffer.WriteString("\n")
			buffer.WriteString("|")
		}

		if cell.IsOccupied() {
			buffer.WriteString(cell.String())
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
	if b.hasHorizontalWin(piece) || b.hasVerticalWin(piece) || b.hasDiagonalWin(piece) {
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

// Checks the board for a diagonal win condition for the provided piece
func (b *Board) hasDiagonalWin(piece string) bool {
	for i := 0; i < (b.Width * b.Height); i += 1 {
		if b.Cells[i].IsOccupiedBy(piece) != true {
			continue
		}

		if b.getNumMatchesTopLeft(i, piece)+b.getNumMatchesBottomRight(i, piece)+1 >= b.MatchesToWin {
			return true
		}

		if b.getNumMatchesTopRight(i, piece)+b.getNumMatchesBottomLeft(i, piece)+1 >= b.MatchesToWin {
			return true
		}
	}

	return false
}

// Gets the number of matches in the bottom left direction
func (b *Board) getNumMatchesBottomLeft(cell int, piece string) int {
	index := cell
	matches := 0

	// While the cell we're checking is not on the top row or left col
	for !(b.isCellOnBottomRow(index) || b.isCellOnLeftCol(index)) {
		index = index + b.Width - 1

		if b.Cells[index].IsOccupiedBy(piece) {
			matches += 1
		} else {
			break
		}
	}

	return matches
}

// Gets the number of matches in the top left direction
func (b *Board) getNumMatchesTopLeft(cell int, piece string) int {
	index := cell
	matches := 0

	// While the cell we're checking is not on the top row or left col
	for !(b.isCellOnTopRow(index) || b.isCellOnLeftCol(index)) {
		index = index - b.Width - 1

		if b.Cells[index].IsOccupiedBy(piece) {
			matches += 1
		} else {
			break
		}
	}

	return matches
}

// Gets the number of matches in the top right direction
func (b *Board) getNumMatchesTopRight(cell int, piece string) int {
	index := cell
	matches := 0

	// While the cell we're checking is not on the top row or left col
	for !(b.isCellOnTopRow(index) || b.isCellOnRightCol(index)) {
		index = index - b.Width + 1

		if b.Cells[index].IsOccupiedBy(piece) {
			matches += 1
		} else {
			break
		}
	}

	return matches
}

// Gets the number of matches in the bottom right direction
func (b *Board) getNumMatchesBottomRight(cell int, piece string) int {
	index := cell
	matches := 0

	// While the cell we're checking is not on the top row or left col
	for !(b.isCellOnBottomRow(index) || b.isCellOnRightCol(index)) {
		index = index + b.Width + 1

		if b.Cells[index].IsOccupiedBy(piece) {
			matches += 1
		} else {
			break
		}
	}

	return matches
}

// Checks whether the cell is on the top border of the board
func (b *Board) isCellOnTopRow(cell int) bool {
	return cell < b.Width
}

// Checks whether the cell is on the bottom border of the board
func (b *Board) isCellOnBottomRow(cell int) bool {
	return (cell >= (b.Height-1)*b.Width)
}

// Checks whether the cell is on the right border of the board
func (b *Board) isCellOnRightCol(cell int) bool {
	return cell%b.Width == (b.Width - 1)
}

// Checks whether the cell is on the left border of the board
func (b *Board) isCellOnLeftCol(cell int) bool {
	return cell%b.Width == 0
}

// Checks whether the board is completely occupied (a tie)
func (b *Board) isFull() bool {
	for _, cell := range b.Cells {
		if !cell.IsOccupied() {
			return false
		}
	}
	return true
}
