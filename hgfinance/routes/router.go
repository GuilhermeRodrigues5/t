package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/hgfinance/controllers"
)

func GetCoin(c *gin.Engine) {
	c.GET("/dollar", controllers.Dollar)
}

func GetBitCoin(c *gin.Engine) {
	c.GET("/bitcoin", controllers.BitCoin)
}
