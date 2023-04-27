package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/hgfinance/routes"
)

func main() {
	router := gin.New()

	routes.GetCoin(router)
	routes.GetBitCoin(router)

	router.Run(":5000")

}
