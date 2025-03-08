package db

import (
	"TikTok/pkg/constants"
	"TikTok/pkg/crypt"
	"context"
	"errors"
)

func CreateUser(ctx context.Context, username, password string) error {
	var userResp *User
	err := DB.WithContext(ctx).
		Table(constants.TableUser).
		Where("BINARY username = ?", username).
		First(&userResp).
		Error
	if err == nil { //找到了
		return errors.New("username already exists")
	}
	psw, err := crypt.PasswordHash(password)
	if err != nil {
		return err
	}
	userResp = &User{
		Username: username,
		Password: psw,
	}
	err = DB.WithContext(ctx).
		Table(constants.TableUser).
		Create(userResp).
		Error
	if err != nil {
		return err
	}
	return nil
}

func LoginCheck(ctx context.Context, username, password string) (*User, error) {
	var userResp *User
	err := DB.WithContext(ctx).
		Table(constants.TableUser).
		Where("BINARY username = ?", username).
		First(&userResp).
		Error
	if err != nil {
		return nil, errors.New("username not found")
	}
	if !crypt.VerifyPassword(password, userResp.Password) {
		return nil, errors.New("password not match")
	} else {
		return userResp, nil
	}
}

func UploadAvatar(ctx context.Context, userid int64, avatar_url string) (*User, error) {
	var userResp *User
	err := DB.WithContext(ctx).
		Table(constants.TableUser).
		Where("id = ?", userid).
		Update("avatar_url", avatar_url).
		Error
	if err != nil {
		return nil, err
	}
	err = DB.WithContext(ctx).
		Table(constants.TableUser).
		Where("id = ?", userid).
		Find(&userResp).
		Error
	if err != nil {
		return nil, err
	}
	return userResp, nil
}

func GetUserInformation(ctx context.Context, userid int64) (*User, error) {
	var userResp *User
	err := DB.WithContext(ctx).
		Table(constants.TableUser).
		Where("id = ?", userid).
		First(&userResp).
		Error
	if err != nil {
		return nil, err
	}
	return userResp, nil
}
