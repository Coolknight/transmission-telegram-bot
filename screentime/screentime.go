package screentime

import (
	"encoding/gob"
	"fmt"
	"os"
)

const logFilePrefix = "config/accountability_"

type Operation struct {
	Description string
	Minutes     int
}

type Accountability struct {
	TotalMinutes int
	Operations   []Operation
}

func Initialize(kidName string) error {
	logFile := logFileName(kidName)

	// Remove existing file if it exists
	_ = os.Remove(logFile)

	// Create a new file to store data
	file, err := os.Create(logFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize an empty Accountability struct
	accountability := Accountability{
		TotalMinutes: 0,
		Operations:   make([]Operation, 0),
	}

	// Encode and write the struct to the file
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(accountability)
	if err != nil {
		return err
	}

	return nil
}

func AddMinutes(kidName string, description string, minutes int) error {
	logFile := logFileName(kidName)

	// Read the existing data
	accountability, err := readData(logFile)
	if err != nil {
		return err
	}

	// Add minutes and update the total
	accountability.TotalMinutes += minutes

	// Add the operation to the log
	accountability.addOperation(description, minutes)

	// Write the updated data back to the file
	return writeData(logFile, accountability)
}

func SubtractMinutes(kidName string, description string, minutes int) error {
	logFile := logFileName(kidName)

	// Read the existing data
	accountability, err := readData(logFile)
	if err != nil {
		return err
	}

	// Subtract minutes and update the total
	accountability.TotalMinutes -= minutes

	// Add the operation to the log
	accountability.addOperation(description, -minutes)

	// Write the updated data back to the file
	return writeData(logFile, accountability)
}

func GetAccountability(kidName string) (string, error) {
	logFile := logFileName(kidName)

	// Read the existing data
	accountability, err := readData(logFile)
	if err != nil {
		return "", err
	}

	// Format the accountability information
	result := fmt.Sprintf("Total Minutes: %d\nOperations:\n", accountability.TotalMinutes)
	for _, op := range accountability.Operations {
		result += fmt.Sprintf("- %s (%+d minutes)\n", op.Description, op.Minutes)
	}

	return result, nil
}

func (a *Accountability) addOperation(description string, minutes int) {
	// Append the new operation to the log, keeping the last 10 operations
	if len(a.Operations) >= 10 {
		a.Operations = a.Operations[1:]
	}
	a.Operations = append(a.Operations, Operation{Description: description, Minutes: minutes})
}

// build the file name
func logFileName(kidName string) string {
	return fmt.Sprintf("%s%s.gob", logFilePrefix, kidName)
}

// Read data from file and return it
func readData(logFile string) (Accountability, error) {
	file, err := os.Open(logFile)
	if err != nil {
		return Accountability{}, err
	}
	defer file.Close()

	var accountability Accountability
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&accountability)
	if err != nil {
		return Accountability{}, err
	}

	return accountability, nil
}

// Write data back to the file.
func writeData(logFile string, accountability Accountability) error {
	file, err := os.Create(logFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(accountability)
	if err != nil {
		return err
	}

	return nil
}
