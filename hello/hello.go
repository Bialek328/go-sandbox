package main

import (
	// "example.com/greetings"
	"fmt"
	// "log"
)

func printPyramid(rows int) {
	spaces_nr := rows - 1
	for i := 1; i <= rows; i++ {
		for j := 1; j <= spaces_nr; j++ {
			fmt.Print(" ")

		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		spaces_nr -= 1
		fmt.Println()
	}
}

func printSquare(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func printRectangle(rows, columns int) {
	str := "* "
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			fmt.Print(str)
		}
		fmt.Println()
	}
}

func chooseFigure() {
	var figure string
	var rows int
	fmt.Println("Choose one figure: rectangle, square, pyramid:")
	_, err1 := fmt.Scanln(&figure)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("How many rows do you want the figure to have?:")
	_, err2 := fmt.Scanln(&rows)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	switch figure {
	case "rectangle":
		var columns int
		fmt.Println("How many columns do you want the figure to have?")
		_, err := fmt.Scanln(&columns)
		if err != nil {
			fmt.Println(err)
			return
		}
		printRectangle(rows, columns)
	case "pyramid":
		printPyramid(rows)
	case "square":
		printSquare(rows)
	default:
		fmt.Println("Wrong shape type")
	}
}
func main() {

	chooseFigure()
}
