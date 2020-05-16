package domain

type ReportItem struct {
	Site                       string
	ResponseTimeInMilliseconds int
	StatusCode                 int
	Error                      string
	TimeFromStartToFirstByte   int
}

type MeasuringResult struct {
	Site                       string
	ResponseTimeInMilliseconds int
	StatusCode                 int
	Error                      string
	TimeFromStartToFirstByte   int
}

type ParsedResult struct {
	Sites []string
}
