package stat

import (
	"errors"
	"site_analyzer/domain"
)

const (
	TestSite1 = "site1"
	TestSite2 = "site2"
)

type ParserMock struct {
	Result     *domain.ParsedResult
	Err        error
	PassedData string
}

func (p *ParserMock) Parse(data string) (*domain.ParsedResult, error) {
	p.PassedData = data
	return p.Result, p.Err
}

func (p *ParserMock) ReturnsSuccess() {
	sites := []string{TestSite1, TestSite2}
	p.Result = &domain.ParsedResult{Sites: sites}
}

func (p *ParserMock) ReturnsError() {
	p.Err = errors.New("some error")
}
