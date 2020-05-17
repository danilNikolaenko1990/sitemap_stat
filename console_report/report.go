package console_report

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"sitemap_stat/domain"
	"sitemap_stat/measure"
	"sitemap_stat/net"
	"sitemap_stat/parsing"
	"sitemap_stat/reporting"
	"sitemap_stat/stat"
	"strconv"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
}

func Report(sitemapUrl string, reportName string, workers int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fetcher := net.NewFetcher()
	measurer := measure.NewMeasurer()
	parser := parsing.NewParser()
	//I dont want to add any di container for 3 dependencies only=)
	analyzer := stat.NewAnalyzer(fetcher, measurer, parser)

	results, tasksCount, err := analyzer.Analyze(sitemapUrl, workers)
	fmt.Println()
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
	waitUntilBarRenderingEnds()
	log.Infof("writing reporting to file %s", reportName)
	err = reporting.Csv(reportName, rep)
	if err != nil {
		log.Errorf("error wile writing reporting %s", err.Error())
	}
	log.Infof("file %s generated", reportName)
}

func waitUntilBarRenderingEnds() {
	//ugly crutch because of slow progress bar rendering
	time.Sleep(1 * time.Second)
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

func createReport() [][]string {
	return [][]string{
		{"Page", "response time in milliseconds", "http status", "time from start to first byte", "error"},
	}
}
