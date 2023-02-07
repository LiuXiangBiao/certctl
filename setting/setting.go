package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() (err error) {

	viper.SetConfigFile("./conf/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("配置读取失败", zap.Error(err))
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Warn("配置文件修改了")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper unmarshal faield", zap.Error(err))
		}
	})
	return
}

var Conf = new(ToolConfig)

type ToolConfig struct {
	Domains_file_path        string  `mapstructure:"domains_file_path"`
	Certbot_config_file_path string  `mapstructure:"certbot_config_file_path"`
	Distance_day_time        float64 `mapstructure:"distance_day_time "`
	AccessKeyId              string  `mapstructure:"accessKeyId"`
	AccessKeySecret          string  `mapstructure:"accessKeySecret"`
	Token                    string  `mapstructure:"token"`
	Secret                   string  `mapstructure:"secret"`
}
