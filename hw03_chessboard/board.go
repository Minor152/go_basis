package main

import (
	"fmt"
	"log"
	"strings"
)

type CHESSBOARD struct {
	board   strings.Builder
	rows    int
	columns int
}

func (b *CHESSBOARD) Init() *CHESSBOARD {
	var err error

	fmt.Println("Enter the number of rows: ")
	_, err = fmt.Scanf("%d", &b.rows)
	if err != nil {
		log.Fatalln(fmt.Errorf("%w", err))
	}

	fmt.Println("Enter the number of columns: ")
	_, err = fmt.Scanf("%d", &b.columns)
	if err != nil {
		log.Fatalln(fmt.Errorf("%w", err))
	}

	return b
}

func (b *CHESSBOARD) Print() string {
	return b.board.String()
}

func (b *CHESSBOARD) Generate() {

	for l := 0; l < b.rows; l++ {
		if l%2 == 0 {

			b.board.WriteString("\n")
			for c := 0; c < b.columns; c++ {
				if c%2 == 0 {
					b.board.WriteString(" ")
				} else {
					b.board.WriteString("#")
				}
			}
		} else {

			b.board.WriteString("\n")
			for c := 0; c < b.columns; c++ {
				if c%2 == 0 {
					b.board.WriteString("#")
				} else {
					b.board.WriteString(" ")
				}
			}
		}

	}
}
