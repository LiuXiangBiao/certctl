package individually

import (
	"crypto/tls"
	"fmt"
	"time"
)

func CheckCertificate(domains []string) {
	for _, domain := range domains {
		//连接服务器获取信息
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), nil)
		if err != nil {
			fmt.Printf("获取%s证书信息失败", domain)
		}

		endTime := conn.ConnectionState().PeerCertificates[0].NotAfter
		fmt.Printf("%s证书到期时间：%d\n", domain, endTime)

		intervalTime := endTime.Sub(time.Now()).Hours() / 24
		fmt.Printf("%s证书到期时间距离今日还有：%f天\n", domain, intervalTime)
	}
}
