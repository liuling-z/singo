package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 视频投稿服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=300"`
}

// setSession 设置session
func (service *CreateVideoService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "视频创建失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
