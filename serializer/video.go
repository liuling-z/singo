package serializer

import "singo/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// File 文件序列化器
type File struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	Src       string `json:"src"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}

// BuildFile 序列化文件
func BuildFile(item model.File) File {
	return File{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		Src:       item.Src,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
