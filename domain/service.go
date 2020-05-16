package domain

type Fetcher interface {
	GetBodyAsString(url string) (string, error)
}

type Measurer interface {
	Measure(url string) MeasuringResult
}

type Parser interface {
	Parse(data string) (*ParsedResult, error)
}
