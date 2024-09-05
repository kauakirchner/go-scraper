package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/go-scraper/cmd/controller"
	"github.com/kauakirchner/go-scraper/cmd/controller/routes"
	"github.com/kauakirchner/go-scraper/cmd/service"
)

func main() {
	router := gin.Default()
	scraperServices := service.NewScraperInterface()

	scraperController := controller.NewInterfaceScraperController(scraperServices)
	routes.InitRoutes(&router.RouterGroup, scraperController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
