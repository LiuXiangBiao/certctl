package automation

import (
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func ReplaceDomains(redomains []string) []string {
	var domain string
	var domains = make([]string, 30, 60)
	for _, redomain := range redomains {
		restr := strings.HasPrefix(redomain, "*")
		if restr == true {
			domain = strings.Replace(redomain, "*", "a", 1)
		} else {
			domain = redomain
		}
		domains = append(domains, domain)
	}
	return domains
}

func FindDirPath(domain string) string {
	domainStr := Prefix(domain)
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("find /etc/letsencrypt/live/ -name '%s*'", domainStr))
	out, err := cmd.StdoutPipe()
	if err != nil {
		zap.L().Error("find domain faield", zap.Error(err))
		return ""
	}
	if err := cmd.Start(); err != nil {
		zap.L().Error("start find cmd faield", zap.Error(err))
		return ""
	}
	res, err := ioutil.ReadAll(out)
	if err != nil {
		zap.L().Error("readall find cmd out faield", zap.Error(err))
		return ""
	}
	return Siffix(res)
}

func Prefix(domain string) string {
	str := strings.HasPrefix(domain, "*.")
	var domainStr string
	if str == true {
		domainStr = strings.Trim(domain, "*.")
	} else {
		return domain
	}
	return domainStr
}

func Siffix(res []byte) string {
	strDir := strings.HasSuffix(string(res), "\n")
	var dirPath string
	if strDir == true {
		dirPath = strings.Trim(string(res), "\n")
	} else {
		return string(res)
	}
	return dirPath
}

func Checkpath(path string, domain string) {
	contain := strings.Contains(path, domain)
	if contain == true {
		zap.L().Info("success get certfile path")
	} else {
		zap.L().Error("certfile path faield")
		os.Exit(2)
	}
}

func GetFileInfo(filepath string) (string, error) {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		zap.L().Error("openfile domains faield", zap.Error(err))
		return "", err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		zap.L().Error("readall domainsfile faield", zap.Error(err))
		return "", err
	}

	return string(data), nil
}
