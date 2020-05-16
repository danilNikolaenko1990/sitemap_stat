package main

import (
	"github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"site_analyzer/domain"
	"site_analyzer/measure"
	"site_analyzer/net"
	"site_analyzer/parsing"
	"site_analyzer/reporting"
	"site_analyzer/stat"
	"strconv"
)

const (
	Url            = ""
	OutputFileName = "reporting.csv"
)

var workers = 100 //todo get from env or commandline

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fetcher := net.NewFetcher()
	measurer := measure.NewMeasurer()
	parser := parsing.NewParser()

	analyzer := stat.NewAnalyzer(fetcher, measurer, parser)

	results, tasksCount, err := analyzer.Analyze(Url, workers)
	if err != nil {
		log.Errorf("error analyzing %s", err.Error())
	}
	bar := createProgressBar(tasksCount)

	rep := createReport()
	for result := range results {
		row := createReportRow(result)
		rep = append(rep, row)
		bar.Increment()
	}
	log.Infof("writing reporting to file %s", OutputFileName)
	err = reporting.Csv(OutputFileName, rep)
	if err != nil {
		log.Errorf("error wile writing reporting %s", err.Error())
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
