package automation

import (
	"certificate/setting"
	"crypto/tls"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"go.uber.org/zap"
	"os"
	"time"
)

func CheckDomainInfo(cfg *setting.ToolConfig, domains []string) {
	//遍历证书并监控到期时间
	//连接获取证书信息
	for _, domain := range domains {
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain), nil)
		if err != nil {
			zap.L().Error("tls dial faield", zap.Error(err))
			return
		}

		//证书到期时间
		endTime := conn.ConnectionState().PeerCertificates[0].NotAfter
		//今天距离证书到期还有几天
		intervalTime := endTime.Sub(time.Now()).Hours() / 24
		zap.L().Warn(fmt.Sprintf("%s endtime=%v intervaltime=%v", domain, endTime, intervalTime))

		var estimatedTime = cfg.Distance_day_time
		if intervalTime < estimatedTime {
			zap.L().Warn(fmt.Sprintf("%s domain intervalTime < estimatedTime start generate certificate", domain))
			//证书到期使用工具生成证书
			certpath, keypath, err := BronCertificate(setting.Conf,domain)
			if err != nil {
				zap.L().Error(fmt.Sprintf("bron certificate %s faield", domain), zap.Error(err))
				break
			} else {
				//上传证书
				if domain == "erda.cloud" {
					certname := domain
					if err := PassToAliYun(setting.Conf, certname, certpath, keypath); err != nil {
						zap.L().Error("upload certificate faield", zap.Error(err))
						os.Exit(2)
					} else {
						zap.L().Info("success upload certificate")
						DingTalk(setting.Conf,domain)
					}
				} else {
					certname := fmt.Sprintf("%s-%v", domain, time.Now().Format("0102"))
					if err := PassToAliYun(setting.Conf, certname, certpath, keypath); err != nil {
						zap.L().Error("upload certificate faield", zap.Error(err))
					} else {
						zap.L().Info("success upload certificate")
						DingTalk(setting.Conf,domain)
					}
				}
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func DingTalk(cfg *setting.ToolConfig, domain string) {
	d := dingtalk.InitDingTalkWithSecret(cfg.Token, cfg.Secret)
	mdmsg := []string{
		"### 警告信息:",
		fmt.Sprintf("- 域名:%s", domain),
		"- 状态：已上传",
		fmt.Sprintf("- 日期:%v", time.Now().Format("2006-01-02")),
		"#### 请及时处理部署相关证书！",
	}
	mobiles := []string{"某某"}
	err := d.SendMarkDownMessageBySlice("test1", mdmsg, dingtalk.WithAtAll(), dingtalk.WithAtMobiles(mobiles))
	if err != nil {
		zap.L().Error("dingding send message faield", zap.Error(err))
		return
	}
}
