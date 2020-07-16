package main

import (
	"fmt"
	"os"
	"os/exec"
)

var field = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

var win bool
var players = []string{"X", "O"}
var turnNumber int

const Size = 3

func main() {

	player := players[0]
	var number string
	for turnNumber < Size*Size && !win {
		fmt.Println("Tic-Tac-Goe by Hennadii")
		turnNumber++
		PrintField()
		fmt.Println("Player is ", player)
		fmt.Print("Enter a number to place: ")
		fmt.Scanln(&number)
		field = PlacePlayer(player, number, field)
		player = ChangePlayer(player)
		ClearScreen()

	}

	fmt.Println("")
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
