package db

import (
	"TikTok/pkg/constants"
	"context"
	"errors"
)

func FollowUser(ctx context.Context, to_userid, action_type, userid int64) error {
	////0关注 1取关
	var FollowResp *UserFollow
	var err error
	if action_type == 0 {
		err = DB.WithContext(ctx).
			Table(constants.TableFollower).
			Where("follower_id = ? AND followee_id = ?", userid, to_userid).
			First(&FollowResp).
			Error
		if err == nil {
			return errors.New("repeated follow!")
		}
		FollowResp = &UserFollow{
			FolloweeID: to_userid,
			FollowerID: userid,
		}
		err = DB.WithContext(ctx).
			Table(constants.TableFollower).
			Create(&FollowResp).
			Error
		if err != nil {
			return err
		}
		return nil
	} else if action_type == 1 {
		err = DB.WithContext(ctx).
			Table(constants.TableFollower).
			Where("follower_id = ? AND followee_id = ?", userid, to_userid).
			First(&FollowResp).
			Error
		if err != nil {
			return errors.New("You never follow!")
		}
		err = DB.WithContext(ctx).
			Table(constants.TableFollower).
			Where("follower_id = ? AND followee_id = ?", userid, to_userid).
			Delete(&FollowResp).
			Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid action_type")
}

// 查看关注列表
func QueryFollowList(ctx context.Context, userid, pagesize, pagenum int64) ([]*User, int64, error) {
	var FollowResp []*UserFollow
	var userResp []*User
	var count int64

	err := DB.WithContext(ctx).
		Table(constants.TableFollower).
		Where("follower_id = ? ", userid).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&FollowResp).
		Error
	if err != nil {
		return nil, -1, err
	}

	//根据获取的关注列表的followee_id查询用户信息
	// 提取被关注者用户id列表
	var to_userIDs []int64
	for _, userId := range FollowResp {
		to_userIDs = append(to_userIDs, userId.FolloweeID)
	}
	// 批量查询用户信息
	if len(to_userIDs) > 0 {
		err = DB.WithContext(ctx).
			Table(constants.TableUser).
			Where("id IN ?", to_userIDs).
			Find(&userResp).
			Error
		if err != nil {
			return nil, -1, err
		}
	}
	return userResp, count, nil
}

// 查看粉丝列表
func QueryFollowerList(ctx context.Context, userid, pagesize, pagenum int64) ([]*User, int64, error) {
	var FollowResp []*UserFollow
	var userResp []*User
	var count int64

	err := DB.WithContext(ctx).
		Table(constants.TableFollower).
		Where("followee_id = ? ", userid).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&FollowResp).
		Error
	if err != nil {
		return nil, -1, err
	}

	//根据获取的关注列表的follower_id查询用户信息
	// 提取粉丝用户id列表
	var followerIDs []int64
	for _, userId := range FollowResp {
		followerIDs = append(followerIDs, userId.FollowerID)
	}
	// 批量查询用户信息
	if len(followerIDs) > 0 {
		err = DB.WithContext(ctx).
			Table(constants.TableUser).
			Where("id IN ?", followerIDs).
			Find(&userResp).
			Error
		if err != nil {
			return nil, -1, err
		}
	}
	return userResp, count, nil
}

// 查看好友列表
func QueryFriendList(ctx context.Context, userid, pagesize, pagenum int64) ([]*User, int64, error) {
	var FollowMEResp []*User
	var MyFollowResp []*User
	var count int64
	var err error

	// 查询我关注的用户
	MyFollowResp, _, err = QueryFollowList(ctx, userid, pagesize, pagenum)
	if err != nil {
		return nil, -1, err
	}

	// 查询关注我的用户
	FollowMEResp, _, err = QueryFollowerList(ctx, userid, pagesize, pagenum)
	if err != nil {
		return nil, -1, err
	}

	// 将 FollowMEResp 转换为 map 以便快速查找
	followMeMap := make(map[int64]*User)
	for _, user := range FollowMEResp {
		followMeMap[user.Id] = user
	}

	// 提取交集（即互相关注的用户）
	var friends []*User
	for _, user := range MyFollowResp {
		if _, exists := followMeMap[user.Id]; exists {
			friends = append(friends, user)
		}
	}

	// 计算总数
	count = int64(len(friends))

	return friends, count, nil
}
