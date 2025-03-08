package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Video struct {
	Id           int64
	UserId       int64
	VideoUrl     string
	CoverUrl     string
	Title        string
	Description  string
	VisitCount   int64
	LikeCount    int64
	CommentCount int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Comment struct {
	Id        int64
	UserId    int64
	Content   string
	VideoId   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserFollow struct {
	FollowerID int64
	FolloweeID int64
	FollowedAt time.Time
}

type VideoLike struct {
	UserId  int64
	VideoId int64
	LikedAt time.Time
}
