package individually

import (
	"fmt"
	"os/exec"
)

func GenerateCertificate(configPath string, domain string) {

	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("certbot certonly -a certbot-dns-aliyun:dns-aliyun  -d %s  --certbot-dns-aliyun:dns-aliyun-credentials %s --preferred-challenges dns-01   --register-unsafely-without-email   --server https://acme-v02.api.letsencrypt.org/directory --force-renewal", domain, configPath))
	//获取标准或者其他输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

}
