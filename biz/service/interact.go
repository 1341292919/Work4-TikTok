package service

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/interact"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type InteractService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewInteractService(ctx context.Context, c *app.RequestContext) *InteractService {
	return &InteractService{ctx: ctx, c: c}
}

func (s *InteractService) LikeVideo(req *interact.LikeVideoRequest) error {
	return db.LikeVideo(s.ctx, GetUserIDFromContext(s.c), req.VideoID, req.ActionType)
}
func (s *InteractService) QueryLikeList(req *interact.QueryLikeListRequest) ([]*db.Video, int64, error) {
	return db.QueryLikeList(s.ctx, req.UserID, req.PageSize, req.PageNum)
}
func (s *InteractService) CommentVideo(req *interact.CommentRequest) error {
	return db.CommentVideo(s.ctx, GetUserIDFromContext(s.c), req.VideoID, req.Content)
}
func (s *InteractService) QueryCommentList(req *interact.QueryCommentListRequest) ([]*db.Comment, int64, error) {
	return db.QueryVideoCommentList(s.ctx, req.VideoID, req.PageSize, req.PageNum)
}
func (s *InteractService) DeleteComment(req *interact.DeleteCommentRequest) error {
	return db.DeleteVideoComment(s.ctx, GetUserIDFromContext(s.c), req.VideoID, req.CommentID)
}
