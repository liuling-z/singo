package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// ListVideoService 视频列表
type ListVideoService struct {
}

// setSession 设置session
func (service *ListVideoService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// ShowList 视频列表
func (service *ListVideoService) ShowList() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Data:  nil,
			Msg:   "视频列表查询失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideos(videos)}
}
