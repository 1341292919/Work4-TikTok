package db

import (
	"TikTok/pkg/constants"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

func LikeVideo(ctx context.Context, userid, videoid, action_type int64) error {
	var LikeResp *VideoLike
	var err error
	//先在点赞列表中查找是否已经点赞过了
	//1 点赞   2 取消点赞
	if action_type == 1 {
		err = DB.
			WithContext(ctx).
			Table(constants.TableLike).
			Where("user_id = ? AND video_id = ?", userid, videoid).
			First(&LikeResp).
			Error
		if err == nil {
			return errors.New("You have hitten the like botton already")
		}

		LikeResp = &VideoLike{
			UserId:  userid,
			VideoId: videoid,
		}
		err = DB.
			WithContext(ctx).
			Table(constants.TableLike).
			Create(&LikeResp).
			Error
		if err != nil {
			return err
		}
		//并在video表中将对应id的视频的like_num加一
		err = DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", videoid).
			Update("like_count", gorm.Expr("like_count + 1")).
			Error
		//DB.Transaction()
		if err != nil {
			return err
		}
		return nil

	} else if action_type == 2 {
		err = DB.WithContext(ctx).
			Table(constants.TableLike).
			Where("user_id = ? AND video_id = ?", userid, videoid).
			First(&LikeResp).
			Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("You nerver like the video")
			}
			return err
		}
		err = DB.WithContext(ctx).
			Table(constants.TableLike).
			Where("user_id = ? AND video_id = ?", userid, videoid).
			Delete(&LikeResp).
			Error
		if err != nil {
			return err
		}
		err = DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", videoid).
			Update("like_count", gorm.Expr("like_count - 1")).
			Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unknown action type")

}

func QueryLikeList(ctx context.Context, userid, pagesize, pagenum int64) ([]*Video, int64, error) {
	var videoResp []*Video
	var count int64
	var likeResp []*VideoLike

	// 查询用户点赞的视频 ID 列表（分页）
	err := DB.WithContext(ctx).
		Table(constants.TableLike).
		Where("user_id = ?", userid).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Find(&likeResp).
		Error
	if err != nil {
		return nil, -1, err
	}

	// 获取点赞总数
	err = DB.WithContext(ctx).
		Table(constants.TableLike).
		Where("user_id = ?", userid).
		Count(&count).
		Error
	if err != nil {
		return nil, -1, err
	}

	// 提取视频 ID 列表
	var videoIDs []int64
	for _, like := range likeResp {
		videoIDs = append(videoIDs, like.VideoId)
	}

	// 批量查询视频信息
	if len(videoIDs) > 0 {
		err = DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id IN ?", videoIDs).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, err
		}
	}
	for _, video := range videoResp {
		err = DB.WithContext(ctx).
			Table(constants.TableVideo).
			Where("id = ?", video.Id).
			Update("visit_count", gorm.Expr("visit_count + 1")).
			Error
		if err != nil {

			return nil, -1, err
		}
	}

	return videoResp, count, nil
}

func CommentVideo(ctx context.Context, userid, videoid int64, content string) error {
	var commentResp *Comment
	var videoResp *Video

	err := DB.WithContext(ctx).
		Table(constants.TableVideo).
		Where("id = ?", videoid).
		First(&videoResp).
		Error
	if err != nil {
		return errors.New("the video_id not exist")
	} //找不到的情况

	err = DB.WithContext(ctx).
		Table(constants.TableVideo).
		Where("id = ?", videoResp.Id).
		Update("comment_count", gorm.Expr("comment_count + 1")).
		Error
	if err != nil {
		// 如果更新失败，记录日志或返回错误
		log.Printf("Failed to update visit_count for video %d: %v", videoResp.Id, err)
		return err
	}

	commentResp = &Comment{
		UserId:  userid,
		VideoId: videoid,
		Content: content,
	}
	err = DB.WithContext(ctx).
		Table(constants.TableComment).
		Create(commentResp).
		Error
	if err != nil {
		return err
	}
	return nil
}

func QueryVideoCommentList(ctx context.Context, videoid, pagesize, pagenum int64) ([]*Comment, int64, error) {
	var commentResp []*Comment
	var count int64
	err := DB.WithContext(ctx).
		Table(constants.TableComment).
		Where("video_id = ?", videoid).
		Where("deleted_at IS NULL"). // 排除已软删除的记录
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&commentResp).
		Error
	if err != nil {
		return nil, -1, err
	}
	return commentResp, count, nil
}

func DeleteVideoComment(ctx context.Context, userid, videoid, commentid int64) error {
	var err error
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if videoid != 0 {
			//删除某个视下的所有评论
			var commentCount int64
			err = tx.Table(constants.TableComment).
				Where("video_id = ? AND user_id = ?", videoid, userid).
				Count(&commentCount).
				Error
			if commentCount == 0 {
				return errors.New("comment under the video not exist")
			}
			if err != nil {
				return err
			}
			err = tx.Table(constants.TableComment).
				Where("video_id = ? AND user_id = ?", videoid, userid).
				Delete(&Comment{}).
				Error
			if err != nil {
				return err
			}
		} else {
			//删除某条特定评论
			var commentResp *Comment
			err = tx.Table(constants.TableComment).
				Where("id = ? AND user_id = ?", commentid, userid).
				First(&commentResp).
				Error
			if err != nil {
				return err
			}
			err = tx.Table(constants.TableComment).
				Where("id = ? AND user_id = ?", commentid, userid).
				Delete(&Comment{}).
				Error
			if err != nil {
				return err
			}
			//更新视频的评论总数
			err = updateCommentCount(tx, videoid, -1)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func updateCommentCount(tx *gorm.DB, videoid int64, delta int) error {
	return tx.Table(constants.TableVideo).
		Where("id = ?", videoid).
		Update("comment_count", gorm.Expr("comment_count + ?", delta)).
		Error
}
