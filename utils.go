package sudoku_nan

import (
	"io/ioutil"
	"net/http"
	"fmt"
)



func IsValid(board [9][9]int, targetX, targetY, value int) bool {
	if value < 0 && 9 < value {
		return false
	}

	for xBoard, boardValue := range board[targetY] {
		if boardValue == value && xBoard != targetX {
			return false
		}
	}

	for yBoard, row := range board {
		if row[targetX] == value && yBoard != targetY {
			return false
		}
	}

	boxX, boxY := (targetX / 3)*3, (targetY / 3)*3
	for yBoard := boxY; yBoard < boxY + 3; yBoard++ {
		for xBoard := boxX; xBoard < boxX + 3; xBoard++ {
			if xBoard != targetX && board[yBoard][xBoard] == value {
				return false
			}
		}
	}

	return true
}


func SudokuChecker(board [9][9]int) bool {
	for y, row := range board {
		for x := range row {
			if !IsValid(board, x, y, board[y][x]) {
				return false
			}
		}
	}
	return true
}


func GetBoardFromWebsite(difficulty string) string {
	resp, err := http.Get("https://sudoku.com/api/level/" + difficulty)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}
	return string(body)
}

func PrintBoard(board [9][9]int) {
	for _, row := range board {
		for _, value := range row {
			fmt.Printf(string(rune(value + 48)) + " ")
		}
		fmt.Println("")
	}
}
