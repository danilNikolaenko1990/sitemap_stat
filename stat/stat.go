package stat

import (
	log "github.com/sirupsen/logrus"
	"sitemap_stat/domain"
	"sync"
)

type Analyzer struct {
	fetcher  domain.Fetcher
	measurer domain.Measurer
	parser   domain.Parser
}

func NewAnalyzer(
	fetcher domain.Fetcher,
	measurer domain.Measurer,
	parser domain.Parser) *Analyzer {
	return &Analyzer{fetcher: fetcher, measurer: measurer, parser: parser}
}

func (a *Analyzer) Analyze(sitemapUrl string, workers int) (result chan domain.ReportItem, tasksCount int, err error) {
	urls, err := a.getUrls(sitemapUrl)
	if err != nil {
		log.Errorf("error while getting urls: %s", err)
		return nil, tasksCount, err
	}
	tasksCount = len(urls)
	tasks := a.queue(urls)
	log.Infof("starting %d workers to measure pages", workers)
	wg := &sync.WaitGroup{}
	results := a.performTasks(wg, tasks, workers)
	go func() {
		wg.Wait()
		close(results)
	}()

	return results, tasksCount, nil
}

func (a *Analyzer) performTasks(wg *sync.WaitGroup, tasks chan string, workers int) chan domain.ReportItem {
	results := make(chan domain.ReportItem, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				results <- a.measurePage(task)
			}
		}()
	}
	return results
}

func (a *Analyzer) measurePage(url string) domain.ReportItem {
	measured := a.measurer.Measure(url)
	return domain.ReportItem{
		ResponseTimeInMilliseconds: measured.ResponseTimeInMilliseconds,
		StatusCode:                 measured.StatusCode,
		Site:                       url,
		Error:                      measured.Error,
		TimeFromStartToFirstByte:   measured.TimeFromStartToFirstByte,
	}

}

func (a *Analyzer) queue(urls []string) chan string {
	jobs := make(chan string)
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	return jobs
}

func (a *Analyzer) getUrls(url string) ([]string, error) {
	xml, err := a.fetcher.GetBodyAsString(url)
	if err != nil {
		return nil, err
	}
	parsed, err := a.parser.Parse(xml)
	if err != nil {
		return nil, err
	}

	return parsed.Sites, nil
}
