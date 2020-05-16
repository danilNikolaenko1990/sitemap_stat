package stat

import (
	"site_analyzer/domain"
)

const (
	TestSite                       = "site"
	TestResponseTimeInMilliseconds = 1
	TestStatusCode                 = 200
	TestError                      = "error"
	TestTimeFromStartToFirstByte   = 2
)

type MeasurerMock struct {
	RecordedPassedData []string
}

func (m *MeasurerMock) Measure(url string) domain.MeasuringResult {
	m.RecordedPassedData = append(m.RecordedPassedData, url)
	return domain.MeasuringResult{
		Site:                       TestSite,
		ResponseTimeInMilliseconds: TestResponseTimeInMilliseconds,
		StatusCode:                 TestStatusCode,
		Error:                      TestError,
		TimeFromStartToFirstByte:   TestTimeFromStartToFirstByte,
	}
}
