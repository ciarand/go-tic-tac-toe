package main

import (
	"fmt"
	"github.com/ciarand/go-tic-tac-toe/lib"
)

func main() {
	fmt.Println("Hello, welcome to tic-tac-toe!")
	var input int
	var board *lib.Board
	var width, height, matches_to_win, number_of_cells int

	fmt.Printf("Enter the width of the board: ")
	fmt.Scanf("%d", &width)

	fmt.Printf("Enter the height of the board: ")
	fmt.Scanf("%d", &height)

	fmt.Printf("Finally, enter the number of matches required to win: ")
	fmt.Scanf("%d", &matches_to_win)

	number_of_cells = width * height

	for board = lib.NewBoard(height, width, matches_to_win); board.IsGameOver() != true; {
		fmt.Println()
		fmt.Println(board)
		fmt.Printf("Enter your move (1-%d)", number_of_cells)
		fmt.Scanf("%d", &input)
		if input > number_of_cells || input < 1 {
			fmt.Println("Out of bounds! Try again")
			continue
		}
		fmt.Printf("Your move was %d", input)

		board.PlacePiece("X", input-1)
	}

	if len(board.Winner) != 0 {
		fmt.Printf("\n%s won the game!", board.Winner)
	} else {
		fmt.Println("It was a tie")
	}

	fmt.Println()
	fmt.Println(board)
}
