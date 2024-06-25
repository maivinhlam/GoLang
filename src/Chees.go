package src

import (
	"fmt"
	"log"
	"strings"
)

type MoveAble struct {
	x int
	y int
}

const (
	row = 9
	col = 9
)

func chees() {
	canMove("E5")
}

func canMove(input string) {
	input = strings.ToUpper(input)
	chars := []rune(input)
	if len(chars) != 2 {
		log.Println("Error")
		return
	}

	positionX := string(chars[0])
	positionY := string(chars[1])

	board := [][]string{}
	for i := 0; i < row; i++ {
		a := []string{}
		for j := 0; j < col; j++ {
			a = append(a, "")
		}
		board = append(board, a)
	}

	for i := 0; i < 8; i++ {
		board[i][0] = fmt.Sprint((i + 1 - row) * -1)
	}

	i := 1
	for c := 'A'; c <= 'Z'; c++ {
		board[8][i] = string(c)
		i++
		if i > 8 {
			break
		}
	}

	px := 0
	py := 0

	for i, v := range board {
		if v[0] == positionY {
			for j, vv := range board[8] {
				if vv == positionX {
					px = i
					py = j
					board[i][j] = "🏇"
					break
				}
			}
		}
	}

	if px == 0 && py == 0 {
		log.Println("out of chees")
		return
	}
	movesAble := []MoveAble{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	for _, v := range movesAble {
		x := px + v.x
		y := py + v.y

		if x > 0 && x < 8 && y > 0 && y < 9 {
			board[x][y] = "♟️"
		}
	}

	for i, v := range board {
		for j := range v {
			if board[i][j] == "♟️" {
				fmt.Printf("%s - %s\n", strings.ToUpper(board[8][j]), board[i][0])
			}
		}
	}
	PrintBoard(board)
}

func PrintBoard(board [][]string) {
	for _, v := range board {
		for _, vv := range v {
			if vv == "" {
				fmt.Print(vv + "- ")
				continue
			}
			fmt.Print(vv + " ")
		}
		fmt.Print("\n")
	}
}
