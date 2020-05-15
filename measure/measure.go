package measure

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

type Result struct {
	Site                       string
	ResponseTimeInMilliseconds int
	StatusCode                 int
	Error                      string
	TimeFromStartToFirstByte   int
}

func Measure(url string) Result {
	req, _ := http.NewRequest("GET", url, nil)
	var start, connect, dns, tlsHandshake time.Time
	result := Result{}
	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		ConnectStart: func(network, addr string) { connect = time.Now() },
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
