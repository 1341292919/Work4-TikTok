package service

import (
	"TikTok/pkg/constants"
	"TikTok/pkg/oss"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"
)

func UploadAvatarGetUrl(data *multipart.FileHeader, userID int64) (string, error) {
	var err error
	err = oss.IsImage(data)
	if err != nil {
		return "", err
	}

	avatarFilename := fmt.Sprintf("%d_avatar", userID)
	avatarPath := filepath.Join(constants.AvatarStorePath, avatarFilename)

	if err = oss.SaveFile(data, constants.AvatarStorePath, avatarFilename); err != nil {
		return "", fmt.Errorf("failed to save avatar: %w", err)
	}
	return avatarPath, nil
}

func UploadVideoGetUrl(data *multipart.FileHeader, userID int64) (string, string, error) {
	var err error

	// 1. 校验是否为视频文件
	if err = oss.IsVideo(data); err != nil {
		return "", "", err
	}

	// 2. 生成视频文件名和路径
	timestamp := time.Now().Unix() // 获取当前时间戳
	videoFilename := fmt.Sprintf("%d_%d_video", userID, timestamp)
	videoPath := filepath.Join(constants.VideoStorePath, videoFilename)

	// 3. 保存视频文件
	if err = oss.SaveFile(data, constants.VideoStorePath, videoFilename); err != nil {
		return "", "", fmt.Errorf("failed to save video: %w", err)
	}

	// 4. 生成封面文件名和路径
	coverFilename := fmt.Sprintf("%d_%d_cover.jpg", userID, timestamp)
	coverPath := filepath.Join(constants.CoverStorePath, coverFilename)

	// 6. 提取视频第一帧作为封面
	if err = oss.ExtractFirstFrame(videoPath, coverPath); err != nil {
		return "", "", errors.New("get cover failed")
	}
	// 7. 返回封面 URL 和视频 URL
	return videoPath, coverPath, nil
}
