package csv_process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type EthereumPrice struct {
	Date  time.Time
	Price float64
}

func LoadDataFrom(path string) ([]EthereumPrice, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	var parsedPairs []EthereumPrice

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read record: %w", err)
		}

		parsedDate, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %w", err)
		}

		parsedPrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price: %w", err)
		}

		parsedPairs = append(parsedPairs, EthereumPrice{
			Date:  parsedDate,
			Price: parsedPrice,
		})
	}

	return parsedPairs, nil
}
