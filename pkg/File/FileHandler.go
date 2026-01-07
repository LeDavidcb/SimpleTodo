package file

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"

	"github.com/SimpleTodo/pkg"
)

func SaveData(todos pkg.TodoList) error {
	// open the file
	file, err := os.Create("save.stodo")
	defer file.Close()
	if err != nil {
		return errors.New("Error while creating/opening file")
	}
	var encoder *gob.Encoder = gob.NewEncoder(file)

	if err := encoder.Encode(&todos); err != nil {
		return fmt.Errorf("Error while encoding struct", err)
	}
	return nil
}
func LoadData(todos *pkg.TodoList) error {
	// Attempt to open the file
	file, err := os.Open("save.stodo")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// If file does not exist, create it
			file, err = os.Create("save.stodo")
			if err != nil {
				return errors.New("failed to create file: " + err.Error())
			}
			fmt.Println("File 'save.stodo' created successfully.")
			file.Close() // Close immediately after creating
			return nil   // Return since there's nothing to load
		}
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	// Decode contents
	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(todos); err != nil {
		return errors.New("error decoding file: " + err.Error())
	}
	return nil
}
