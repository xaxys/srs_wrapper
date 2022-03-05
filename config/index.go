package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	AppConfig *viper.Viper
)

func init() {
	AppConfig = viper.New()
	AppConfig.SetConfigName("appconfig")
	AppConfig.SetConfigType("yaml")
	AppConfig.AddConfigPath(".")
	AppConfig.AddConfigPath("./config")
	AppConfig.AddConfigPath("/etc/srs_wrappper/")
	AppConfig.AddConfigPath("$HOME/.srs_wrappper/")

	AppConfig.SetDefault("app.name", "srs_wrapper")
	AppConfig.SetDefault("app.listen", ":8787")
	AppConfig.SetDefault("app.loglevel", "info")
	AppConfig.SetDefault("app.jwtkey", "srs_wrapper_2022_all_rights_reserved")

	AppConfig.SetDefault("database.driver", "sqlite")
	AppConfig.SetDefault("database.sqlite.path", "srs_wrapper.db")
	AppConfig.SetDefault("database.mysql.host", "localhost")
	AppConfig.SetDefault("database.mysql.port", "3306")
	AppConfig.SetDefault("database.mysql.name", "srs_wrapper")
	AppConfig.SetDefault("database.mysql.params", "parseTime=true&loc=Local&charset=utf8mb4")
	AppConfig.SetDefault("database.mysql.user", "root")
	AppConfig.SetDefault("database.mysql.password", "123456")

	AppConfig.SetDefault("cache.expire", 86400)
	AppConfig.SetDefault("cache.purge", 600)

	AppConfig.SetDefault("admin.name", "admin")
	AppConfig.SetDefault("admin.display_name", "srs_wrapper default admin")
	AppConfig.SetDefault("admin.password", "123456")

	if err := AppConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			AppConfig.SafeWriteConfig()
		} else {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}
}
