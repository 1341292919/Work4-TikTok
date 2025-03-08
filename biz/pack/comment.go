package pack

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/model"
	"strconv"
)

func Comment(data *db.Comment) *model.Comment {
	return &model.Comment{
		ID:        data.Id,
		UserID:    data.UserId,
		VideoID:   data.VideoId,
		Content:   data.Content,
		CreatedAt: strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdatedAt: strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}

func CommentList(data []*db.Comment, total int64) *model.CommentList {
	resp := make([]*model.Comment, 0, len(data))
	for _, v := range data {
		resp = append(resp, Comment(v))
	}
	return &model.CommentList{
		Items: resp,
		Total: total,
	}
}
