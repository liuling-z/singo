package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// LoadFile 文件上传
func LoadFile(c *gin.Context) {
	service := service.LoadFileService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.LoadFile()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
