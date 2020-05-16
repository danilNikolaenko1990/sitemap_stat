package stat

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sitemap_stat/domain"
	"testing"
)

const Url = "testurl"
const Worker = 1

func TestAnalyze_GivenSiteMapUrl_ReturnsResultToChannel(t *testing.T) {
	fetcherMock := createFetcherWithSuccessResult()
	measurerMock := &MeasurerMock{}
	parserMock := createParserWithSuccess()
	analyzer := NewAnalyzer(fetcherMock, measurerMock, parserMock)
	results, tasksCount, err := analyzer.Analyze(Url, Worker)
	assert.Nil(t, err)
	assert.Equal(t, len(parserMock.Result.Sites), tasksCount)

	countOfResults := 0
	for result := range results {
		countOfResults++
		assertResultItem(t, result)
	}
	assert.Equal(t, 2, countOfResults)
	assert.Equal(t, fetcherMock.Res, parserMock.PassedData)
	assert.Equal(t, TestSite1, measurerMock.RecordedPassedData[0])
	assert.Equal(t, TestSite2, measurerMock.RecordedPassedData[1])
}

func TestAnalyze_GivenFetcherReturnsError_AnalyzerReturnsError(t *testing.T) {
	fetcherMock := &FetcherMock{}
	fetcherMock.Err = errors.New("some err")

	measurerMock := &MeasurerMock{}
	parserMock := createParserWithSuccess()
	analyzer := NewAnalyzer(fetcherMock, measurerMock, parserMock)
	_, _, err := analyzer.Analyze(Url, Worker)
	assert.EqualError(t, err, "some err")
}

func TestAnalyze_GivenParserReturnsError_AnalyzerReturnsError(t *testing.T) {
	fetcherMock := createFetcherWithSuccessResult()
	measurerMock := &MeasurerMock{}
	parserMock := &ParserMock{}
	parserMock.Err = errors.New("some err")
	analyzer := NewAnalyzer(fetcherMock, measurerMock, parserMock)
	_, _, err := analyzer.Analyze(Url, Worker)
	assert.EqualError(t, err, "some err")
}

func assertResultItem(t *testing.T, item domain.ReportItem) {
	assert.Equal(t, TestResponseTimeInMilliseconds, item.ResponseTimeInMilliseconds)
	assert.Equal(t, TestStatusCode, item.StatusCode)
	assert.Equal(t, TestTimeFromStartToFirstByte, item.TimeFromStartToFirstByte)
	assert.Equal(t, TestError, item.Error)
	assert.NotEmpty(t, item.Site)
}

func createFetcherWithSuccessResult() *FetcherMock {
	fetcher := &FetcherMock{}
	fetcher.Res = "some res"

	return fetcher
}

func createParserWithSuccess() *ParserMock {
	parser := &ParserMock{}
	parser.ReturnsSuccess()
	return parser
}
