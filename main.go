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

const Size = 3

func main() {
	fmt.Println("Tic-Tac-Goe by Hennadii")

	PrintField()
	fmt.Scanln()
	fmt.Println("")
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
