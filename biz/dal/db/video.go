package db

import (
	"TikTok/pkg/constants"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

func UploadVideo(ctx context.Context, video_url, cover_url, title, description string, userid int64) error {
	var videoResp *Video
	//要不要检验一下文件路径？
	//标题相同会出现什么情况
	videoResp = &Video{
		Title:       title,
		Description: description,
		VideoUrl:    video_url,
		CoverUrl:    cover_url,
		UserId:      userid,
	}
	err := DB.WithContext(ctx).
		Table(constants.TableVideo).
		Create(&videoResp).
		Error
	if err != nil {
		return err
	}
	return nil
}

// 查询发布列表
func QueryVideoList(ctx context.Context, userid, pagesize, pagenum int64) ([]*Video, int64, error) {
	var videoResp []*Video
	var count int64

	// 查询视频列表
	err := DB.WithContext(ctx).
		Table(constants.TableVideo).
		Where("user_id = ?", userid).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&videoResp).
		Error
	if err != nil {
		return nil, -1, err
	}

	// 遍历查询到的视频，将每个视频的 visit_count 加一
	for _, video := range videoResp {
		err := DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", video.Id).
			Update("visit_count", gorm.Expr("visit_count + 1")).
			Error
		if err != nil {
			// 如果更新失败，记录日志或返回错误
			log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
			return nil, -1, err
		}
	}

	return videoResp, count, nil
}

func SearchVideoByKeywordDuringTime(ctx context.Context, keyword string, pagesize, pagenum int64, to_date, from_date time.Time) ([]*Video, int64, error) {
	var videoResp []*Video
	var count int64
	// 关键词模糊匹配
	keyword = "%" + keyword + "%"
	err := DB.WithContext(ctx).
		Table(constants.TableVideo).
		Where("created_at >= ? AND created_at <= ? ", from_date, to_date).
		Where("title LIKE ? OR description LIKE ?", keyword, keyword).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&videoResp).
		Error
	if err != nil {
		return nil, -1, err
	}
	// 遍历查询到的视频，将每个视频的 visit_count 加一
	for _, video := range videoResp {
		err := DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", video.Id).
			Update("visit_count", gorm.Expr("visit_count + 1")).
			Error
		if err != nil {
			// 如果更新失败，记录日志或返回错误
			log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
			return nil, -1, err
		}
	}
	return videoResp, count, nil
}

func SearchVideoByKeywordDuringTimeAndUser(ctx context.Context, username, keyword string, pagesize, pagenum int64, to_date, from_date time.Time) ([]*Video, int64, error) {
	var videoResp []*Video
	var user *User
	var count int64
	// 关键词模糊匹配
	var err error

	err = DB.WithContext(ctx).
		Table(constants.TableUser).
		Where("username = ?", username).
		Find(&user).
		Error
	if err != nil {
		return nil, -1, errors.New("username not found")
	}

	keyword = "%" + keyword + "%"
	err = DB.WithContext(ctx).
		Table(constants.TableVideo).
		Where("user_id = ?", user.Id).
		Where("created_at >= ? AND created_at <= ? ", from_date, to_date).
		Where("title LIKE ? OR description LIKE ?", keyword, keyword).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&videoResp).
		Error
	if err != nil {
		return nil, -1, err
	}
	// 遍历查询到的视频，将每个视频的 visit_count 加一
	for _, video := range videoResp {
		err := DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", video.Id).
			Update("visit_count", gorm.Expr("visit_count + 1")).
			Error
		if err != nil {
			// 如果更新失败，记录日志或返回错误
			log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
			return nil, -1, err
		}
	}
	return videoResp, count, nil
}

func QueryVideoByPopularity(ctx context.Context, pagesize, pagenum int64) ([]*Video, int64, error) {
	var videoResp []*Video
	var count int64

	err := DB.WithContext(ctx).
		Table(constants.TableVideo).
		Order("visit_count DESC").
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&videoResp).
		Error
	if err != nil {
		return nil, -1, err
	}

	// 遍历查询到的视频，将每个视频的 visit_count 加一
	for _, video := range videoResp {
		err := DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", video.Id).
			Update("visit_count", gorm.Expr("visit_count + 1")).
			Error
		if err != nil {
			// 如果更新失败，记录日志或返回错误
			log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
			return nil, -1, err
		}
	}
	return videoResp, count, nil
}
