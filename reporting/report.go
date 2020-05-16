package reporting

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Csv(filename string, dataToWrite [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error while creating file %w", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	for _, record := range dataToWrite {
		if err := w.Write(record); err != nil {
			return fmt.Errorf("error while recording %w", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("error while flushing dataToWrite %w", err)
	}
	return nil
}
