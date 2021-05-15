package main

import (
	"fmt"
	"go-playground/common"
	"go-playground/config"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// 测试用的入口文件
func main() {
	/*f, err := os.Open("user-service/user-service")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		panic(err)
	}

	contentType := http.DetectContentType(buffer)
	fmt.Println(contentType)*/

	sftpClient, err := common.CreateSftp(
		config.SSH_KEY,
		config.SSH_USER,
		config.SSH_HOST,
		config.SSH_PORT,
	)
	if err != nil {
		return
	}
	defer sftpClient.Close()

	// 上传文件
	localFilePath := "./ssh_key.pem"      // 本地文件路径
	remoteDir := config.REMOTE_UPLOAD_DIR // 远程目录

	srcFile, err := os.Open(localFilePath)
	fileBytes, err := ioutil.ReadAll(srcFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	remoteFilename := path.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFilename))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	dstFile.Write(fileBytes)
	fmt.Println("文件上传完成")
}
