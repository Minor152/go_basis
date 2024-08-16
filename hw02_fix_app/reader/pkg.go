package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Minor152/go_basis/hw02_fix_app/types"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Println(err)
	}

	res := data

	return res, nil
}
