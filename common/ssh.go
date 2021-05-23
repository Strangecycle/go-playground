package common

import (
	"fmt"
	"github.com/pkg/sftp"
	"go-playground/config"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"
)

// 使用密钥连接到 ssh
func ConnectSSHWithKey(sshKey string, sshUser string, sshHost string, sshPort int) (*ssh.Client, error) {
	// 读取密钥
	key, err := ioutil.ReadFile(sshKey)
	if err != nil {
		panic("Fail to read the SSH key")
	}

	// 解析密钥签名
	privateKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic("Fail to sign the SSH key")
	}

	// ssh 配置
	sshConfig := &ssh.ClientConfig{
		User:    sshUser,
		Timeout: 5 * time.Second,
		Auth: []ssh.AuthMethod{
			// 使用密钥登录远程服务器
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这个不够安全，生产环境不建议使用
		// HostKeyCallback: ssh.FixedHostKey(), //建议使用这种
	}

	// 连接远程服务器
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		fmt.Println("Fail to connect remote consumer：" + err.Error())
		return nil, err
	}
	return client, nil
}

// 创建 sftp 会话，传输文件
func CreateSftp(sshKey string, sshUser string, sshHost string, sshPort int) (*sftp.Client, error) {
	conn, err := ConnectSSHWithKey(sshKey, sshUser, sshHost, sshPort)
	if err != nil {
		panic("Fail to connect remote consumer：" + err.Error())
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// 获取 sftp 连接客户端
func GetSftpClient() (*sftp.Client, error) {
	sftpClient, err := CreateSftp(
		// TODO 为什么在 linux 上只能读当前目录下的文件？？
		config.SshKey,
		config.SshUser,
		config.SshHost,
		config.SshPort,
	)
	if err != nil {
		panic("Fail to get sftp client：" + err.Error())
	}

	return sftpClient, nil
}
