package scraper

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
)

type Anime struct {
	Name string
	Link string
}

const searchPrefix = "https://www.animeworld.ac/search?keyword="

func SearchAnime(query string) (results []Anime) {
	if len(query) == 0 {
		fmt.Println("Can't lookup empty query")
		return nil
	}
	query = strings.Replace(query, " ", "+", -1)
	var searchUrl string = searchPrefix + query

	results = []Anime{}

	c := colly.NewCollector()

	c.OnHTML("a[href].name", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		name := e.Attr("data-jtitle")

		found := Anime{
			Name: name,
			Link: e.Request.AbsoluteURL(link),
		}

		results = append(results, found)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(searchUrl)

	return results
}

func GetDownloadInfo(animeMainPageLink string) (epCount int, downloadPageLink string)  {
	c := colly.NewCollector()

	/* [TODO] some longer anime have double the episode count */
	c.OnHTML("li.episode", func(e *colly.HTMLElement) {
		epCount++
	})

	c.OnHTML("#downloadLink", func(e *colly.HTMLElement) { 
		downloadPageLink = e.Attr("href")
		downloadPageLink = e.Request.AbsoluteURL(downloadPageLink)
	})

	c.Visit(animeMainPageLink)

	return epCount, downloadPageLink
}

func GetDirectDlLink(downloadPageLink string) (directDlLink string) {
	c := colly.NewCollector()

	c.OnHTML("a.btn", func(e *colly.HTMLElement) {
		directDlLink = e.Attr("href")
		directDlLink = e.Request.AbsoluteURL(directDlLink)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.Visit(downloadPageLink)

	return directDlLink 
}
