package parser

import "encoding/xml"

type (
	urlSet struct {
		XMLName xml.Name `xml:"urlset"`
		Urls    []url    `xml:"url"`
	}

	url struct {
		XMLName    xml.Name `xml:"url"`
		Loc        string   `xml:"loc"`
		Priority   string   `xml:"priority"`
		Changefreq string   `xml:"changefreq"`
	}

	sitemapindex struct {
		XMLName  xml.Name  `xml:"sitemapindex"`
		Sitemaps []sitemap `xml:"sitemap"`
	}

	sitemap struct {
		XMLName xml.Name `xml:"sitemap"`
		Loc     string   `xml:"loc"`
	}
)
