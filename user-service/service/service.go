package service

import "go-playground/proto/user"

func AddUser(phone string) user.AddUserResponse {
	return user.AddUserResponse{
		Message: "success",
	}
}
