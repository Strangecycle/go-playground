package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/logger"
	"go-playground/common"
	"go-playground/proto/user"
	"go-playground/user-service/model"
	"go-playground/util"
)

// SendCaptcha
func SendCaptcha(request *user.CaptchaRequest) user.CaptchaResponse {
	// TODO 将验证码存进 redis 为登录时做验证
	conn := common.GetRedisConnect()
	defer conn.Close()

	captcha := util.GenCaptcha()
	conn.Do("Set", "captcha", captcha)
	conn.Do("expire", "captcha", 120) // 120 秒后过期

	return user.CaptchaResponse{
		Captcha: captcha,
	}
}

// UserLogin
func UserLogin(phone string, captcha string) user.UserLoginResponse {
	response := user.UserLoginResponse{Token: ""}
	if phone == "" || captcha == "" {
		return response
	}

	conn := common.GetRedisConnect()
	defer conn.Close()

	// 对比验证码
	code, _ := redis.String(conn.Do("Get", "captcha"))
	if code != captcha {
		return response
	}

	db := common.GetDB()
	userModel := &model.User{
		Username: "", // 默认用户名为手机号
		Phone:    phone,
		Avatar:   "",
		Email:    "",
		Sentence: "",
	}
	db.Where("phone = ?", phone).First(&userModel)
	// 如果用户不存在则创建用户
	if userModel.ID == 0 {
		if err := db.Create(&userModel).Error; err != nil {
			logger.Fatal(err.Error())
		}
	}

	response.Token = common.GenToken(userModel.Username, userModel.ID)

	// 登录成功后验证码失效
	conn.Do("Del", "captcha")

	return response
}

// UserInfo
func UserInfo(phone string, email string) user.UserInfoResponse {
	db := common.GetDB()
	userModel := &model.User{}

	if phone != "" {
		db.Where("phone = ?", phone).First(&userModel)
	}
	if email != "" {
		db.Where("email = ?", email).First(&userModel)
	}

	if userModel.ID == 0 {
		return user.UserInfoResponse{}
	}

	return user.UserInfoResponse{
		Id:        uint64(userModel.ID),
		Username:  userModel.Username,
		Avatar:    userModel.Avatar,
		Phone:     userModel.Phone,
		Email:     userModel.Email,
		Sentence:  userModel.Sentence,
		CreatedAt: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userModel.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
