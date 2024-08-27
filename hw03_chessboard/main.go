package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	// Count colums.
	SizeColumn int
	// Count rows.
	SizeRow int
	// String containing a chessboard.
	CHESSBOARD strings.Builder
)

func init() {
	var err error

	fmt.Println("Enter the number of rows: ")
	_, err = fmt.Scanf("%d", &SizeRow)
	if err != nil {
		log.Fatalln(fmt.Errorf("%w", err))
	}

	fmt.Println("Enter the number of columns: ")
	_, err = fmt.Scanf("%d", &SizeColumn)
	if err != nil {
		log.Fatalln(fmt.Errorf("%w", err))
	}
}

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
