package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService 删除视频信息接口
type DeleteVideoService struct {
}

// setSession 设置session
func (service *DeleteVideoService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Delete 删除视频信息
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.Where("id=?", id).Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "删除视频失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
