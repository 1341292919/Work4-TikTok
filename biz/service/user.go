package service

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"mime/multipart"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}
func (s *UserService) Login(req *user.LoginRequest) (*db.User, error) {
	return db.LoginCheck(s.ctx, req.Username, req.Password)
}
func (s *UserService) Register(req *user.RegisterRequest) error {
	return db.CreateUser(s.ctx, req.Username, req.Password)
}
func (s *UserService) UploadAvatar(data *multipart.FileHeader) (*db.User, error) {
	userid := GetUserIDFromContext(s.c)
	url, err := UploadAvatarGetUrl(data, userid)
	if err != nil {
		return nil, err
	}
	return db.UploadAvatar(s.ctx, userid, url)
}

func (s *UserService) GetUserInformation(req *user.GetUserInformationRequest) (*db.User, error) {
	return db.GetUserInformation(s.ctx, req.UserID)
}
