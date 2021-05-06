package service

import (
	"github.com/micro/go-micro/v2/logger"
	"go-playground/config"
	"go-playground/proto/file"
	"io/ioutil"
)

func SingleUpload(request *file.SingleUploadRequest) file.SingleUploadResponse {
	err := ioutil.WriteFile("./upload/"+request.GetFilename(), request.GetFile(), 0666)
	if err != nil {
		logger.Error(err.Error())
		return file.SingleUploadResponse{}
	}

	return file.SingleUploadResponse{
		Filename: request.GetFilename(),
		FileUrl:  config.DOMAIN + "/upload/" + request.GetFilename(),
	}
}
