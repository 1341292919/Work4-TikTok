package pack

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/model"
	"strconv"
)

func Video(data *db.Video) *model.Video {
	return &model.Video{
		ID:           data.Id,
		Userid:       data.UserId,
		VideoURL:     data.VideoUrl,
		CoverURL:     data.CoverUrl,
		Title:        data.Title,
		Description:  data.Description,
		VisitCount:   data.VisitCount,
		LikeCount:    data.LikeCount,
		CommentCount: data.CommentCount,
		CreatedAt:    strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdatedAt:    strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}

func VideoList(data []*db.Video, total int64) *model.VideoList {
	resp := make([]*model.Video, 0, len(data))
	for _, v := range data {
		resp = append(resp, Video(v))
	}
	return &model.VideoList{
		Items: resp,
		Total: total,
	}
}
