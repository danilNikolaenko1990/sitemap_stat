package parsing

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"site_analyzer/domain"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(data string) (*domain.ParsedResult, error) {
	log.Println("parsing xml")
	urls := urlSet{}
	sitemaps := sitemapindex{}

	if err := xml.Unmarshal([]byte(data), &urls); err == nil {
		sites := make([]string, len(urls.Urls))
		for key, url := range urls.Urls {
			sites[key] = url.Loc
		}

		return &domain.ParsedResult{Sites: sites}, nil
	} else if err := xml.Unmarshal([]byte(data), &sitemaps); err == nil {
		sites := make([]string, len(sitemaps.Sitemaps))
		for key, url := range sitemaps.Sitemaps {
			sites[key] = url.Loc
		}

		return &domain.ParsedResult{Sites: sites}, nil
	} else {
		return nil, err
	}
}
