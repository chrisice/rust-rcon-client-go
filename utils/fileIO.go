package utils

import (
	"bufio"
	"errors"
	"os"
)

// WriteToFile writes the provided data to a file with the given filename.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	return writer.Flush()
}

// ReadFromFile reads data from a file with the given filename.
func ReadFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var content string

	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

// FileExists checks if a file with the given filename exists.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// DeleteFile deletes a file with the given filename.
func DeleteFile(filename string) error {
	if !FileExists(filename) {
		return errors.New("file does not exist")
	}
	return os.Remove(filename)
}
