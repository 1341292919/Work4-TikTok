package service

import (
	"TikTok/biz/dal/db"
	"TikTok/biz/model/video"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"mime/multipart"
	"time"
)

type VideoService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewVideoService(ctx context.Context, c *app.RequestContext) *VideoService {
	return &VideoService{ctx: ctx, c: c}
}

func (s *VideoService) UploadVideo(data *multipart.FileHeader, req *video.PublishRequest) error {
	userid := GetUserIDFromContext(s.c)
	video_url, cover_url, err := UploadVideoGetUrl(data, userid)
	if err != nil {
		return err
	}
	return db.UploadVideo(s.ctx, video_url, cover_url, req.Title, req.Description, userid)
}

func (s *VideoService) QueryPublishedVideo(req *video.QueryPublishListRequest) ([]*db.Video, int64, error) {
	return db.QueryVideoList(s.ctx, req.UserID, req.PageSize, req.PageNum)
}

func (s *VideoService) QueryVideoByKeyword(req *video.SearchVideoByKeywordRequest) ([]*db.Video, int64, error) {

	var fromDate, toDate time.Time
	// 如果 FromDate 未传入（值为 0），设置为最小时间
	if req.FromDate == 0 {
		fromDate = time.Time{} // 最小时间
	} else {
		fromDate = time.Unix(req.FromDate, 0)
	}
	// 如果 ToDate 未传入（值为 0），设置为当前时间
	if req.ToDate == 0 {
		toDate = time.Now()
	} else {
		toDate = time.Unix(req.ToDate, 0)
	}
	if fromDate.Unix() > toDate.Unix() {
		return nil, 0, errors.New("fromDate is after toDate")
	}
	return db.SearchVideoByKeywordDuringTime(s.ctx, req.Keyword, req.PageSize, req.PageNum, toDate, fromDate)
}

func (s *VideoService) QueryVideoByPopularity(req *video.GetPopularListRequest) ([]*db.Video, int64, error) {
	return db.QueryVideoByPopularity(s.ctx, req.PageSize, req.PageNum)
}
