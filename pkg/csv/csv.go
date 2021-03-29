package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

// Read reads the content of the CSV file.
func Read(filePath string) ([][]string, error) {
	csvFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Error().Err(err).Msg("failed to open the csv file")
		return nil, err
	}

	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	var parsedFile [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error().Err(err).Msg("failed to read the csv file")
			return [][]string{}, err
		}

		parsedFile = append(parsedFile, record)
	}

	return parsedFile, err
}

// Write writes the new content to the end of the file.
func Write(filePath string, content string) error {
	csvFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error().Err(err).Msg("could not open csv file to write")
		return err
	}

	defer csvFile.Close()

	if content == "" {
		return errors.New("empty content")
	}

	if _, err := csvFile.WriteString(content + "\n"); err != nil {
		log.Error().Err(err).Msg("failed to write to the csv file")
		return err
	}

	return nil
}
