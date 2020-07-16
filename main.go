package main

import (
	"fmt" // for printing and scanning
	"os"  // for cleaning terminal
	"os/exec"
)

var field = [3][3]string{ // our gameplay field
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

var win bool
var players = []string{"X", "O"}
var turnNumber int
var usedNumbers []string

const Size = 3 // field size

func main() {

	player := players[0]
	var number string
	var winner string

	for turnNumber < Size*Size && !win {
		// remove terminal garbage
		ClearScreen()
		// author signature :))))
		fmt.Println("Tic-Tac-Goe by Hennadii")
		//fmt.Println(turnNumber)
		// drawing field and player
		PrintField()
		fmt.Println("Player is ", player)
		// giving user ability to choose place
		fmt.Print("Enter a number to place: ")
		fmt.Scanln(&number)
		// checking for duplicate player placement
		if !IsNumberUsed(usedNumbers, number) {
			usedNumbers = append(usedNumbers, number)
			turnNumber++
		}
		// Making a turn and changing player
		field = PlacePlayer(player, number, field)
		player = ChangePlayer(player)

		winner, win = CheckWin()

	}

	ClearScreen()
	fmt.Println("Tic-Tac-Goe by Hennadii")
	PrintField()

	if win {
		fmt.Println("Winner is ", winner)
	} else {
		fmt.Println("We got a draw")
	}

	// love fp code snippet
	fmt.Println("")
}

// Checking if number is in slice
func IsNumberUsed(used []string, number string) bool {
	for _, v := range used {
		if number == v {
			return true
		}
	}
	return false
}

func ThreeEqual(a, b, c string) bool {
	return a == b && b == c
}

func CheckWin() (string, bool) {
	u1 := field[0][0]
	u2 := field[0][1]
	u3 := field[0][2]
	u4 := field[1][0]
	u5 := field[1][1]
	u6 := field[1][2]
	u7 := field[2][0]
	u8 := field[2][1]
	u9 := field[2][2]
	switch {
	case ThreeEqual(u1, u2, u3):
		return u1, true
	case ThreeEqual(u4, u5, u6):
		return u4, true
	case ThreeEqual(u7, u8, u9):
		return u7, true
	case ThreeEqual(u1, u4, u7):
		return u1, true
	case ThreeEqual(u2, u5, u8):
		return u2, true
	case ThreeEqual(u3, u6, u9):
		return u3, true
	case ThreeEqual(u1, u5, u9):
		return u1, true
	case ThreeEqual(u3, u5, u7):
		return u3, true
	default:
		return "", false
	}
}

// Places a player
func PlacePlayer(player, number string, field [3][3]string) [3][3]string {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if field[i][j] == number {
				field[i][j] = player
			}
		}
	}
	return field
}

// Changes to X or O
func ChangePlayer(p string) string {
	var index int
	for i := 0; i < len(players); i++ {
		if p == players[i] {
			index = i
		}
	}

	if index == 0 {
		return "O"
	}
	return "X"
}

// Print a game field
func PrintField() {
	fmt.Println("------------")
	for i := 0; i < Size; i++ {

		for j := 0; j < Size; j++ {

			fmt.Print(" ", field[i][j], " |")

			if j == Size-1 {
				fmt.Printf("\n")
			}

		}
		fmt.Println("------------")
	}
}

// clean up terminal
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
