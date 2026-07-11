package array

import (
	"fmt"
	"math/rand"
)

func TicTacToe(size int) {
	board := make([][]string, size)
	for i := range len(board) {
		board[i] = make([]string, size)
	}
	DrawTictacToeBoard(board)
	turnOrder := 1
	playerOne := "1"
	playerTwo := "2"
	for !CheckState(board) {
		if turnOrder == 1 {
			fmt.Println("Player One")
			RandomPlay(playerOne, board)
			turnOrder += 1
		} else if turnOrder == 2 {
			fmt.Println("Player Two")
			RandomPlay(playerTwo, board)
			turnOrder -= 1
		}
		DrawTictacToeBoard(board)
		fmt.Print("\n")
	}
	if turnOrder == 2 {
		fmt.Print("Gagné par le Joueur ", 1)
	} else {
		fmt.Print("Gagné par le Joueur ", 2)
	}
}

func RandomPlay(player string, board [][]string) {
	n := len(board)
	played := false
	for !played {
		x, y := rand.Intn(n), rand.Intn(n)
		if board[x][y] == "" {
			board[x][y] = player
			played = true
		}
	}
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

func CheckState(board [][]string) bool {
	r := false
	for i := range len(board) {
		r = CheckLine(board[i])
		if r {
			return r
		}
		r = CheckColumn(board, i)
		if r {
			return r
		}
	}
	r = CheckDiagonal(board)
	return r
}

func CheckLine(line []string) bool {
	for i := range len(line) {
		val := line[i]
		if val == "" {
			continue
		}
		if i+2 < len(line) {
			if line[i+1] == val && line[i+2] == val {
				return true
			}
		}
	}
	return false
}

func CheckColumn(board [][]string, index int) bool {
	for i := range len(board) {
		val := board[index][i]
		if val == "" {
			continue
		}
		if i+2 < len(board) {
			if board[index][i+1] == val && board[index][i+2] == val {
				return true
			}
		}
	}
	return false
}

func CheckDiagonal(board [][]string) bool {
	n := len(board)
	for i := range n {
		for j := range n {
			val := board[i][j]
			if val == "" {
				continue
			}
			if i+2 < n && j+2 < n {
				if board[i+1][j+1] == val && board[i+2][j+2] == val {
					return true
				}
			}
			if i+2 < n && j-2 >= 0 {
				if board[i+1][j-1] == val && board[i+2][j-2] == val {
					return true
				}
			}
		}
	}
	return false
}
