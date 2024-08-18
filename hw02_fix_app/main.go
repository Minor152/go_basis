package main

import (
	"fmt"
	"log"

	"github.com/Minor152/go_basis/hw02_fix_app/printer"
	"github.com/Minor152/go_basis/hw02_fix_app/reader"
	"github.com/Minor152/go_basis/hw02_fix_app/types"
)

func main() {

	var err error
	var staff []types.Employee

	path := "data.json"

	fmt.Printf("Enter data file path: ")
	if _, err := fmt.Scanln(&path); err != nil {
		fmt.Println(err)
	}

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		log.Fatal(err)
	}

	printer.PrintStaff(staff)
}
