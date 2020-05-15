package main

import (
	"github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"site_analyzer/analyzer"
	"site_analyzer/domain"
	"site_analyzer/report"
	"strconv"
)

const (
	//todo move url to env
	Url            = "https://www.the-village.ru/"
	OutputFileName = "report.csv"
)

var workers = 100 //todo get from env or commandline

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	results, tasksCount := analyzer.Analyze(Url, workers)
	bar := createProgressBar(tasksCount)

	rep := createReport()
	for result := range results {
		row := createReportRow(result)
		rep = append(rep, row)
		bar.Increment()
	}
	log.Infof("writing report to file %s", OutputFileName)
	err := report.Csv(OutputFileName, rep)
	if err != nil {
		log.Errorf("error wile writing report %s", err.Error())
	}
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

func createProgressBar(count int) *pb.ProgressBar {
	bar := pb.StartNew(count)
	return bar
}
