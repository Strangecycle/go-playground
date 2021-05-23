package service

import (
	"github.com/micro/go-micro/v2/logger"
	"go-playground/common"
	"go-playground/config"
	"go-playground/proto/file"
	"path"
)

func SingleUpload(request *file.SingleUploadRequest) file.SingleUploadResponse {
	// 将文件上传到远程服务器
	// fmt.Println("ssh key 文件位置：", path.Join(util.GetParentDir(), config.SSH_KEY))
	sftpClient, _ := common.GetSftpClient()
	defer sftpClient.Close()

	remoteFile, err := sftpClient.Create(path.Join(config.RemoteUploadDir, request.GetFilename()))
	if err != nil {
		logger.Error(err.Error())
		return file.SingleUploadResponse{}
	}
	defer remoteFile.Close()

	remoteFile.Write(request.GetFile())

	return file.SingleUploadResponse{
		Filename: request.GetFilename(),
		FileUrl:  config.FileAddr + request.GetFilename(),
	}

	// err := ioutil.WriteFile("./upload/"+request.GetFilename(), request.GetFile(), 0666)
	// 线上
	// rootPath, _ := os.Getwd()
	// err := ioutil.WriteFile(path.Join(rootPath, "upload", request.GetFilename()), request.GetFile(), 0666)

	// 使用 os 写入文件
	/*openedFile, err := os.OpenFile("./"+request.GetFilename(), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logger.Error(err.Error())
		return file.SingleUploadResponse{}
	}
	defer openedFile.Close()

	writer := bufio.NewWriter(openedFile)
	size, _ := writer.Write(request.GetFile())
	if err != nil {
		logger.Error(err.Error())
		return file.SingleUploadResponse{}
	}

	// 刷新缓冲区，强制写出
	writer.Flush()

	logger.Info("上传文件成功，文件大小为: " + string(size))*/
}
