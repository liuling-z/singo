package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// LoadFileService 上传文件服务
type LoadFileService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=300"`
	Src   string `form:"src" json:"info" binding:"min=0,max=300"`
}

// setSession 设置session
func (service *LoadFileService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// LoadFile 文件上传
func (service *LoadFileService) LoadFile() serializer.Response {
	file := model.File{
		Title: service.Title,
		Info:  service.Info,
		Src:   service.Src,
	}
	err := model.DB.Create(&file).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "文件上传失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildFile(file)}
}
