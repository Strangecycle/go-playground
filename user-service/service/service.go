package service

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/logger"
	"go-playground/common"
	"go-playground/proto/file"
	"go-playground/proto/user"
	"go-playground/user-service/client"
	"go-playground/user-service/model"
	"go-playground/util"
)

// SendCaptcha
func SendCaptcha(request *user.CaptchaRequest) user.CaptchaResponse {
	// TODO request 接收一个手机号
	if request.GetPhone() == "" {
		return user.CaptchaResponse{Captcha: ""}
	}

	conn := common.GetRedisConnect()
	defer conn.Close()

	captcha := util.GenCaptcha()

	// TODO 创建一个短信服务以消息队列（RabbitMQ）的方式给用户发送短信验证码
	go func() {}()

	// 存入 redis
	conn.Do("Set", "captcha", captcha)
	conn.Do("expire", "captcha", 120) // 120 秒后过期

	return user.CaptchaResponse{
		Captcha: captcha,
	}
}

// UserLogin
func UserLogin(phone string, captcha string) user.UserLoginResponse {
	response := user.UserLoginResponse{Token: ""}
	// 为了服务健壮性，在这里也验证一下参数
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
	var userModel *model.User

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

func UserAvatar(request *user.UserAvatarRequest) user.UserAvatarResponse {
	// 调用文件服务上传文件
	fileService := client.GetFileService()
	requestUpload := file.SingleUploadRequest{
		File:     request.GetAvatarBytes(),
		Filename: request.GetAvatarName(),
	}

	responseUpload, err := fileService.SingleUpload(context.Background(), &requestUpload)
	if err != nil {
		logger.Error(err.Error())
		return user.UserAvatarResponse{}
	}

	responseAvatar := user.UserAvatarResponse{
		AvatarName: responseUpload.GetFilename(),
		AvatarUrl:  responseUpload.GetFileUrl(),
	}

	// 修改用户 avatar 字段信息
	db := common.GetDB()
	if err := db.Model(&model.User{}).Where("id = ?", request.GetUserId()).Update("avatar", responseAvatar.GetAvatarUrl()).Error; err != nil {
		logger.Error(err.Error())
	}

	return responseAvatar
}

// UserEdit
func UserEdit(request *user.UserEditRequest) user.UserEditResponse {
	var userModel *model.User
	db := common.GetDB()

	result := db.Model(&userModel).Where("id = ?", request.GetId()).Updates(model.User{
		Username: request.GetUsername(),
		Avatar:   request.GetAvatar(),
		Phone:    request.GetPhone(),
		Email:    request.GetEmail(),
		Sentence: request.GetSentence(),
	})

	if result.Error != nil {
		logger.Info(result.Error.Error())
	}

	return user.UserEditResponse{AffectedRow: result.RowsAffected}
}
