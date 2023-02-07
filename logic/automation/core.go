package automation

import (
	"bufio"
	"certificate/setting"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"io"
	"os"
)

func CronCertificate() {
	// 每天执行一次检查路径下文件内证书到期
	spec := "* * */24 * * *"
	c := cron.New(cron.WithSeconds())
	c.AddFunc(spec, func() {
		var redomains = make([]string, 30, 60)
		var domains = make([]string, 30, 60)

		fin, err := os.OpenFile(setting.Conf.Domains_file_path, os.O_RDONLY, 0666)
		if err != nil {
			zap.L().Error("open domain file faield", zap.Error(err))
			return
		}
		defer fin.Close()

		reader := bufio.NewReader(fin)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			redomains = append(domains, string(line))
		}
		domains = ReplaceDomains(redomains)
		CheckDomainInfo(setting.Conf, domains)
	})
	go c.Start()
	defer c.Stop()
	select {}
}
