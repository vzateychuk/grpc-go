package serializer

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

func WriteToBinaryFile(msg proto.Message, filename string) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message to binary: %w", err)
	}

	// Ensure the directory exists
	dir := filepath.Dir(filename)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write binary data to file: %w", err)
	}

	return nil
}

// ReadFromBinaryFile reads a protocol buffer message from binary file
func ReadFromBinaryFile(filename string, message interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read file: %w", err)
	}

	err = proto.Unmarshal(data, message.(proto.Message))
	if err != nil {
		return fmt.Errorf("unable to unmarshal binary to proto message: %w", err)
	}

	return nil
}

// WriteToJSONFile writes a protocol buffer message to a JSON file
func WriteToJSONFile(message proto.Message, filename string) error {

	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message to JSON: %w", err)
	}

	// Ensure the directory exists
	dir := filepath.Dir(filename)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write JSON data to file: %w", err)
	}

	return nil
}
