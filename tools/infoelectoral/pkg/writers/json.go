package writers

import (
	json2 "encoding/json"
	"fmt"
	"os"
)

func CreateJsonFile(data interface{}, filename string) error {
	var f, err = os.Create(filename + ".json")
	if err != nil {
		return fmt.Errorf("could not create file %s: %w", filename, err)
	}
	defer f.Close()

	var json = json2.NewEncoder(f)
	err = json.Encode(data)
	if err != nil {
		return fmt.Errorf("cannot write json on file %s: %w", filename, err)
	}

	return nil
}
