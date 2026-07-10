package array

import "fmt"

func TicTacToe(size int) {
	board := make([][]string, size)
	for i := range len(board) {
		board[i] = make([]string, size)
	}
	DrawTictacToeBoard(board)
}

func DrawTictacToeBoard(board [][]string) {
	for i := range len(board) {
		fmt.Print("|")
		for j := range len(board) {
			if board[i][j] != "" {
				fmt.Print(board[i][j])
			} else {
				fmt.Print("_")
			}
			fmt.Print("|")
		}
		fmt.Print("\n")
	}
}
