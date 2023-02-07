package individually

import (
	"fmt"
	cas20200407 "github.com/alibabacloud-go/cas-20200407/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io/ioutil"
	"os"
)

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *cas20200407.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("cas.aliyuncs.com")
	_result = &cas20200407.Client{}
	_result, _err = cas20200407.NewClient(config)
	return _result, _err
}

func UploadCertificate(id string, secret string, certName string, certFilePath string, keyFilePath string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(id), tea.String(secret))
	if _err != nil {
		return _err
	}

	cert := GetFileInfo(certFilePath)
	key := GetFileInfo(keyFilePath)

	uploadUserCertificateRequest := &cas20200407.UploadUserCertificateRequest{
		Name: tea.String(certName),
		Cert: tea.String(cert),
		Key:  tea.String(key),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()

		_, _err = client.UploadUserCertificateWithOptions(uploadUserCertificateRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}

		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func GetFileInfo(filepath string) string {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("读取文件失败")
	}

	return string(data)
}
