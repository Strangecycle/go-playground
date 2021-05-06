package handler

import (
	"context"
	"go-playground/file-service/service"
	"go-playground/proto/file"
)

type handler struct {
}

func (h handler) SingleUpload(ctx context.Context, request *file.SingleUploadRequest, response *file.SingleUploadResponse) error {
	*response = service.SingleUpload(request)
	return nil
}

func Handler() file.FileHandler {
	return new(handler)
}
