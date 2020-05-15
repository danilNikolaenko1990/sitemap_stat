package report

import (
	"encoding/csv"
	"fmt"
	"os"
	"site_analyzer/domain"
	"strconv"
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

func createReport() [][]string {
	return [][]string{
		{"Page", "response time in milliseconds", "http status", "time from start to first byte", "error"},
	}
}

func createReportRow(result domain.ReportItem) []string {
	return []string{
		result.Site,
		strconv.Itoa(result.ResponseTimeInMilliseconds),
		strconv.Itoa(result.StatusCode),
		strconv.Itoa(result.TimeFromStartToFirstByte),
		result.Error,
	}
}
