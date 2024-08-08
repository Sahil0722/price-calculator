package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

	if err != nil {
		return nil, errors.New("Error Occured Reading file")
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error Occured Scanning file")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJsonToFile(data any) error {
	file, err := os.Create(fm.OutputPath)

	if err != nil {
		return errors.New("Failed to create file")
	}

	encoder := json.NewEncoder(file) // creating object
	err = encoder.Encode(data)       // calling method on object

	if err != nil {
		return errors.New("Failed to convert data to json")
	}

	file.Close()

	return nil
}
