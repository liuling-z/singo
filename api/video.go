package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateVideo 视屏投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
