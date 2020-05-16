package parsing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_GivenXmlWithXmlLinksInside_ReturnsParsedResult(t *testing.T) {
	sitemapXml := givenSitemapWithSitemapIndexLinks()

	parser := NewParser()
	result, err := parser.Parse(sitemapXml)

	assert.Nil(t, err)
	assert.Equal(t, "https://test.com/sitemap.default.xml", result.Sites[0])
	assert.Equal(t, "https://test2.com/sitemap.default.xml", result.Sites[1])
}

func TestParser_GivenXmlWithUrlSetLinks_ReturnsParsedResult(t *testing.T) {
	sitemapXml := givenSitemapWithUrlSetLinks()

	parser := NewParser()
	result, err := parser.Parse(sitemapXml)

	assert.Nil(t, err)
	assert.Equal(t, "https://test.com/", result.Sites[0])
	assert.Equal(t, "https://test.com/blablabla", result.Sites[1])
}

func TestParser_GivenNotXml_ReturnsError(t *testing.T) {
	someString := `some string`

	parser := NewParser()
	_, err := parser.Parse(someString)

	assert.NotNil(t, err)
}

func givenSitemapWithUrlSetLinks() string {
	return `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<url>
<loc>https://test.com/</loc>
<lastmod>2020-05-16T16:13:43+03:00</lastmod>
<changefreq>daily</changefreq>
<priority>1.0</priority>
</url>
<url>
<loc>https://test.com/blablabla</loc>
<lastmod>2020-05-16T16:13:43+03:00</lastmod>
<changefreq>daily</changefreq>
<priority>0.9</priority>
</url>
</urlset>`
}

func givenSitemapWithSitemapIndexLinks() string {
	return `<sitemapindex xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/siteindex.xsd">
<sitemap>
<loc>https://test.com/sitemap.default.xml</loc>
<lastmod>2020-05-16T16:08:20+03:00</lastmod>
</sitemap>
<sitemap>
<loc>https://test2.com/sitemap.default.xml</loc>
<lastmod>2020-05-16T16:08:20+03:00</lastmod>
</sitemap>
</sitemapindex>`
}
