package middlewares

import (
	"net/http"
	"time"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, utils.ResponseWhenFail("timeout"))
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(30*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
