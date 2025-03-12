package pack

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/model"
	"strconv"
)

func SimpleUser(data *db.User) *model.SimpleUser {
	return &model.SimpleUser{
		ID:        strconv.FormatInt(data.Id, 10),
		Username:  data.Username,
		AvatarURL: data.AvatarUrl,
	}
}
func SimpleUserList(data []*db.User, total int64) *model.SimpleUserList {
	Resp := make([]*model.SimpleUser, 0, len(data))
	for _, o := range data {
		Resp = append(Resp, SimpleUser(o))
	}
	return &model.SimpleUserList{
		Items: Resp,
		Total: total,
	}
}
