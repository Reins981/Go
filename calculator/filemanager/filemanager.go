package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	inputFile  string
	outputFile string
}

func (m FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(m.inputFile)

	if err != nil {
		return nil, errors.New("open file failed")
	}

	defer file.Close() // go will not execute this close operation immediately but when the surrounding method or functions has finished (ReadLines)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("reading the file content failed")
	}

	// file.Close()

	return lines, nil
}

func (m FileManager) WriteResult(data interface{}) error { // interface{} or any -> Accept any value
	file, err := os.Create(m.outputFile)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) // Simulate a slow file writing process

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("failed to convert data to json")
	}

	// file.Close()
	return nil
}

func New(inputFile string, outputFile string) FileManager {
	return FileManager{
		inputFile:  inputFile,
		outputFile: outputFile,
	}
}
