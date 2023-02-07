package automation

import (
	"certificate/setting"
	cas20200407 "github.com/alibabacloud-go/cas-20200407/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
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

// 需要ak，sk
func PassToAliYun(cfg *setting.ToolConfig, certName string, certpath string, keypath string) (_err error) {

	client, _err := CreateClient(tea.String(cfg.AccessKeyId), tea.String(cfg.AccessKeySecret))
	if _err != nil {
		return _err
	}

	// 获取文件路径

	cert, err := GetFileInfo(certpath)
	if err != nil {
		zap.L().Error("getfileinfo faield", zap.Error(err))
		os.Exit(2)
	}
	key, err := GetFileInfo(keypath)
	if err != nil {
		zap.L().Error("getfileinfo faield", zap.Error(err))
		os.Exit(2)
	}

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
