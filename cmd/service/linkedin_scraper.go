package service

import (
	"fmt"
	"slices"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/kauakirchner/go-scraper/pkg/utils"
)

func (i *interfaceScraper) ScrapJobs(
	selectors *Selectors,
	searchUrl string,
) []Job {

	var jobs []Job
	jobs = nil
	start := 0

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utils.RandomString())
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL: ", r.Request.URL, "\nError: ", err)
		panic(err)
	})

	c.OnHTML(selectors.JobSelector, func(e *colly.HTMLElement) {
		jobTitle := e.ChildText(selectors.TitleSelector)
		jobSrc := e.ChildAttr(selectors.SrcSelector, "href")
		company := e.ChildText(selectors.CompanySelector)
		companySrc := e.ChildAttr(selectors.CompanySelector, "href")
		location := e.ChildText(selectors.LocationSelector)
		postedAt := e.ChildAttr(selectors.PostedAtSelector, "datetime")

		foundJob := Job{
			jobTitle,
			jobSrc,
			company,
			companySrc,
			location,
			postedAt,
		}

		isNewJob := !slices.Contains(jobs, foundJob)
		if isNewJob {
			jobs = append(jobs, foundJob)
		}
	})

	for start < 50 {
		c.Visit(fmt.Sprintf("%s%d", searchUrl, start))
		start += 25
		time.Sleep(1 * time.Second)
	}

	return jobs
}
