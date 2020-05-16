package measure

import (
	"net/http"
	"net/http/httptrace"
	"site_analyzer/domain"
	"time"
)

type Measurer struct{}

func NewMeasurer() *Measurer {
	return &Measurer{}
}

func (*Measurer) Measure(url string) domain.MeasuringResult {
	req, _ := http.NewRequest("GET", url, nil)
	var start time.Time
	result := domain.MeasuringResult{}
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			result.TimeFromStartToFirstByte = toMilliseconds(time.Since(start))
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		result.Error = err.Error()
		return result
	}

	result.ResponseTimeInMilliseconds = toMilliseconds(time.Since(start))
	result.StatusCode = resp.StatusCode

	return result
}

func toMilliseconds(duration time.Duration) int {
	return int(duration.Milliseconds())
}
