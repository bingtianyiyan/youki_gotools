package configexternal

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Redis Redis `json:"redis" ini:"redis"`
}

// redis配置模型
type Redis struct {
	Address string `mapstructure:"address" json:"address" ini:"address"`
	Maxidle int `mapstructure:"maxidle" json:"maxidle" ini:"maxidle"`
}

var (
	CONFIG = new(Config)
)

func IniConfigFromYaml()  {
	//动态替换这边文件路径
	//rootCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	var cfgFile string = "../config/config.yaml"
	vtoml := viper.New()
	if cfgFile != "" {
		// Use config file from the flag.
		vtoml.SetConfigFile(cfgFile)
	} else {
		vtoml.AddConfigPath("./")
		vtoml.SetConfigName("config")
		vtoml.SetConfigType("yaml")
	}

	vtoml.AutomaticEnv() // read in environment variables that match

	err := vtoml.ReadInConfig()
	if  err != nil {
		fmt.Println("'config.yaml' file read error:", err)
		os.Exit(0)
	}
	vtoml.Unmarshal(CONFIG)
	fmt.Println("config: ", CONFIG, "redis: ", CONFIG.Redis)
}