package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// UpdateVideoService 视频信息更新服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=300"`
}

// setSession 设置session
func (service *UpdateVideoService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Update 更新视频信息
func (service *UpdateVideoService) Update(id string) serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Model(&video).Update("id", id).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "更新视频失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
