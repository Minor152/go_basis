package main

import "fmt"

func main() {
	var chessboard CHESSBOARD

	chessboard.Init().Generate()

	fmt.Println(chessboard.Print())
}
