package service

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/socialize"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type SocializeService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewSocializeService(ctx context.Context, c *app.RequestContext) *SocializeService {
	return &SocializeService{ctx: ctx, c: c}
}

func (s *SocializeService) FollowUser(req *socialize.FollowRequest) error {
	return db.FollowUser(s.ctx, req.ToUserID, req.ActionType, GetUserIDFromContext(s.c))
}

func (s *SocializeService) QueryFollowList(req *socialize.QueryFollowListRequest) ([]*db.User, int64, error) {
	return db.QueryFollowList(s.ctx, req.UserID, req.PageSize, req.PageNum)
}

// // 查看粉丝列表
func (s *SocializeService) QueryFollowerList(req *socialize.QueryFollowerListRequest) ([]*db.User, int64, error) {
	return db.QueryFollowerList(s.ctx, req.UserID, req.PageSize, req.PageNum)
}
func (s *SocializeService) QueryFriendList(req *socialize.QueryFriendListRequest) ([]*db.User, int64, error) {
	return db.QueryFriendList(s.ctx, GetUserIDFromContext(s.c), req.PageSize, req.PageNum)
}
