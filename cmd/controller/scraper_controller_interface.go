package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/go-scraper/cmd/service"
)

func NewInterfaceScraperController(
	service service.InterfaceScraper,
) InterfaceScraperController {
	return &interfaceScraperController{
		service: service,
	}
}

type InterfaceScraperController interface {
	GetLinkedinJobs(ctx *gin.Context)
}

type interfaceScraperController struct {
	service service.InterfaceScraper
}
