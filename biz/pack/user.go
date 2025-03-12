package pack

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/model"
	"strconv"
)

func User(data *db.User) *model.User {
	return &model.User{
		ID:        strconv.FormatInt(data.Id, 10),
		Username:  data.Username,
		AvatarURL: data.AvatarUrl,
		CreatedAt: strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdatedAt: strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}
