package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Minor152/go_basis/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
