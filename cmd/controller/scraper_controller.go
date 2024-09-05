package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/go-scraper/cmd/service"
)

const (
	SHOW_BODY_ON_ERROR = false

	MAIN_URL     = "https://www.linkedin.com/"
	SEARCH_ROUTE = "jobs/search/"

	JOB_SELECTOR     = ".base-card"
	SRC_SELECTOR     = ".base-card__full-link"
	TITLE_SELECTOR   = "div.base-search-card__info > h3"
	COMPANY_SELECTOR = "div.base-search-card__info > h4 > a"

	LOCATION_SELECTOR  = "div.base-search-card__info > div > span"
	POSTED_AT_SELECTOR = "div.base-search-card__info > div > time"
)

type jobResonse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Jobs    []service.Job `json:"jobs"`
}

func (scraperController *interfaceScraperController) GetLinkedinJobs(ctx *gin.Context) {
	keywords := ctx.Query("keywords")
	excluded := ctx.Query("excluded")
	location := ctx.Query("location")
	remote := ctx.Query("remote")

	searchUrl := fmt.Sprintf(
		"%s%s?location=%s&pageNum=0&keywords=%s NOT \"%s\"&f_WT=%s&start=",
		MAIN_URL, SEARCH_ROUTE, location, keywords, excluded, remote,
	)

	selectors := &service.Selectors{
		TitleSelector:    TITLE_SELECTOR,
		SrcSelector:      SRC_SELECTOR,
		CompanySelector:  COMPANY_SELECTOR,
		LocationSelector: LOCATION_SELECTOR,
		PostedAtSelector: POSTED_AT_SELECTOR,
		JobSelector:      JOB_SELECTOR,
	}

	jobs := scraperController.service.ScrapJobs(selectors, searchUrl)

	response := jobResonse{
		http.StatusOK,
		fmt.Sprintf("%d jobs found!", len(jobs)),
		jobs,
	}

	if len(jobs) > 0 {
		ctx.IndentedJSON(http.StatusOK, response)
		return
	}

	ctx.IndentedJSON(http.StatusBadGateway, "Something wrong happened")
}
