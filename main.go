package main

import (
	"fmt"
	"github.com/ciarand/go-tic-tac-toe/lib"
)

func main() {
	printHeader()
	var board *lib.Board
	width, height, matches_to_win := getBoardSize()

	number_of_cells := width * height

	for board = lib.NewBoard(height, width, matches_to_win); board.IsGameOver() != true; {
		fmt.Println()
		fmt.Println(board)
		input := getNextMove(number_of_cells)
		if input > number_of_cells || input < 1 {
			fmt.Println("Out of bounds! Try again")
			continue
		}
		fmt.Printf("Your move was %d", input)

		board.PlacePiece("X", input-1)
	}

	printResults(board)
}

func printResults(b *lib.Board) {
	if len(b.Winner) != 0 {
		fmt.Printf("\n%s won the game!", b.Winner)
	} else {
		fmt.Println("It was a tie")
	}

	fmt.Println()
	fmt.Println(b)
}

func getBoardSize() (width, height, matches_to_win int) {
	fmt.Printf("Enter the width of the board: ")
	fmt.Scanf("%d", &width)

	fmt.Printf("Enter the height of the board: ")
	fmt.Scanf("%d", &height)

	fmt.Printf("Finally, enter the number of matches required to win: ")
	fmt.Scanf("%d", &matches_to_win)

	return width, height, matches_to_win
}

func printHeader() {
	fmt.Println("Hello, welcome to tic-tac-toe!")
}

func getNextMove(number_of_cells int) (input int) {
	fmt.Printf("Enter your move (1-%d)", number_of_cells)
	fmt.Scanf("%d", &input)
	return input
}
