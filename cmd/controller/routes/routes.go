package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/go-scraper/cmd/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	scraperController controller.InterfaceScraperController,
) {
	r.GET("/scrape", scraperController.GetLinkedinJobs)
}
