package main

import (
	"github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"site_analyzer/fetcher"
	"site_analyzer/measure"
	"site_analyzer/parser"
	"site_analyzer/report"
	"strconv"
	"sync"
)

const (
	//todo move url to env
	Url            = "https://elama.ru/sitemap.blog.xml"
	OutputFileName = "report.csv"
)

var workers = 100 //todo get from env or commandline

func init() {
	log.SetOutput(os.Stdout)
}

type Result struct {
	Site                       string
	ResponseTimeInMilliseconds int
	StatusCode                 int
	Error                      string
	TimeFromStartToFirstByte   int
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	urls, err := getUrls()
	if err != nil {
		panic(err.Error())
	}
	tasks := queue(urls)
	log.Infof("starting %d workers to measure pages", workers)
	wg := &sync.WaitGroup{}
	results := performTasks(wg, tasks)
	go func() {
		wg.Wait()
		close(results)
	}()

	pb := createProgressBar(len(urls))
	log.Info("measuring pages")

	rep := createReport()
	for result := range results {
		row := createReportRow(result)
		rep = append(rep, row)
		pb.Increment()
	}
	log.Infof("writing report to file %s", OutputFileName)
	err = report.Csv(OutputFileName, rep)
	if err != nil {
		log.Errorf("error wile writing report %s", err.Error())
	}
}

func createReportRow(result Result) []string {
	return []string{
		result.Site,
		strconv.Itoa(result.ResponseTimeInMilliseconds),
		strconv.Itoa(result.StatusCode),
		strconv.Itoa(result.TimeFromStartToFirstByte),
		result.Error,
	}
}

func createReport() [][]string {
	return [][]string{
		{"Page", "response time in milliseconds", "http status", "time from start to first byte", "error"},
	}
}

func performTasks(wg *sync.WaitGroup, tasks chan string) chan Result {
	results := make(chan Result, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				results <- measurePage(task)
			}
		}()
	}
	return results
}

func measurePage(url string) Result {
	measured := measure.Measure(url)
	return Result{
		ResponseTimeInMilliseconds: measured.ResponseTimeInMilliseconds,
		StatusCode:                 measured.StatusCode,
		Site:                       url,
		Error:                      measured.Error,
		TimeFromStartToFirstByte:   measured.TimeFromStartToFirstByte,
	}

}

func queue(urls []string) chan string {
	jobs := make(chan string)
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	return jobs
}

func getUrls() ([]string, error) {
	xml, err := fetcher.GetBodyAsString(Url)
	if err != nil {
		return nil, err
	}
	parsed, err := parser.Parse(xml)
	if err != nil {
		return nil, err
	}

	return parsed.Sites, nil
}

func createProgressBar(count int) *pb.ProgressBar {
	bar := pb.StartNew(count)
	return bar
}
