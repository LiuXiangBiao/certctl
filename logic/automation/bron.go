package automation

import (
	"certificate/setting"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os/exec"
)

func BronCertificate(cfg *setting.ToolConfig,domain string) (certpath string, keypath string, err error) {
	// certbot 生成证书
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("certbot certonly -a certbot-dns-aliyun:dns-aliyun  -d %s  --certbot-dns-aliyun:dns-aliyun-credentials %s --preferred-challenges dns-01   --register-unsafely-without-email   --server https://acme-v02.api.letsencrypt.org/directory --force-renewal", domain,cfg.Certbot_config_file_path))
	// 执行Linux命令
	if err = cmd.Run(); err != nil {
		zap.L().Error("certbot certonly cmd run faield", zap.Error(err))
		return
	} else {
		zap.L().Info("success start cmd")
	}

	//检查前缀与后缀过滤出目录路径
	dirPath := FindDirPath(domain)

	//获取fullchainpem证书文件路径
	certcmd := exec.Command("ls", fmt.Sprintf("%s/fullchain.pem", dirPath))
	output1, err := certcmd.StdoutPipe()
	if err != nil {
		zap.L().Error("output full.pem faield", zap.Error(err))
		return "", "", err
	}
	// 读取所有输出
	certPathOut, err := ioutil.ReadAll(output1)
	if err != nil {
		zap.L().Error("ioutil.readall full.pem faield", zap.Error(err))

	}
	certFilePath := Siffix(certPathOut)

	//获取证书私钥文件路径
	keycmd := exec.Command("ls", fmt.Sprintf("%s/privkey.pem", dirPath))
	output2, err := keycmd.StdoutPipe()
	if err != nil {
		zap.L().Error("output key.pem faield", zap.Error(err))
		return "", "", err
	}
	// 读取所有输出
	keyFileOut, err := ioutil.ReadAll(output2)
	if err != nil {
		zap.L().Error("ioutil.readall key.pem faield", zap.Error(err))
		return "", "", err
	}
	keyFilePath := Siffix(keyFileOut)

	Checkpath(certFilePath, domain)
	return certFilePath, keyFilePath, nil
}
