package analyzer

import (
	log "github.com/sirupsen/logrus"
	"runtime"
	"site_analyzer/domain"
	"site_analyzer/fetcher"
	"site_analyzer/measure"
	"site_analyzer/parser"
	"sync"
)

func Analyze(sitemapUrl string, workers int) (result chan domain.ReportItem, tasksCount int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	urls, err := getUrls(sitemapUrl)
	if err != nil {
		log.Fatalf("error while getting urls: %s", err)
		return
	}
	tasksCount = len(urls)
	tasks := queue(urls)
	log.Infof("starting %d workers to measure pages", workers)
	wg := &sync.WaitGroup{}
	results := performTasks(wg, tasks, workers)
	go func() {
		wg.Wait()
		close(results)
	}()

	return results, tasksCount
}

func performTasks(wg *sync.WaitGroup, tasks chan string, workers int) chan domain.ReportItem {
	results := make(chan domain.ReportItem, workers)
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

func measurePage(url string) domain.ReportItem {
	measured := measure.Measure(url)
	return domain.ReportItem{
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

func getUrls(url string) ([]string, error) {
	xml, err := fetcher.GetBodyAsString(url)
	if err != nil {
		return nil, err
	}
	parsed, err := parser.Parse(xml)
	if err != nil {
		return nil, err
	}

	return parsed.Sites, nil
}
