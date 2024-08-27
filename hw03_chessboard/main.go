package main

import (
	"fmt"
)

func main() {
	for l := 0; l < SizeRow; l++ {
		for c := 0; c < SizeColumn; c++ {
			if c == 0 {
				fmt.Fprintf(&CHESSBOARD, "\n")
			}
			fmt.Fprintf(&CHESSBOARD, "# ")
		}
	}
	fmt.Println(CHESSBOARD.String())
}
