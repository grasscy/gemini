package config

import (
	"github.com/Jarnpher553/gemini/log"
	"github.com/Jarnpher553/viper"
)

type Config = viper.Viper

var logger = log.Zap.Mark("config")

//默认使用当前文件夹下的config.yaml文件
//func init() {
//	viper.AddConfigPath(".")
//	viper.SetConfigName("config")
//	viper.SetConfigType("yaml")
//
//	generate()
//}

var conf *Config

func Conf() *Config {
	return conf
}

// Option 配置项方法
type Option func(*viper.Viper)

type Factory func()

// Path 路径配置
func Path(path string) Option {
	return func(v *viper.Viper) {
		v.AddConfigPath(path)
	}
}

// Name 文件名配置
func Name(name string) Option {
	return func(v *viper.Viper) {
		v.SetConfigName(name)
	}
}

// Type 文件类型配置
func Type(_type string) Option {
	return func(v *viper.Viper) {
		v.SetConfigType(_type)
	}
}

func Provider(provider string, endpoint string, keyOrPath string) Option {
	return func(v *viper.Viper) {
		if err := v.AddRemoteProvider(provider, endpoint, keyOrPath); err != nil {
			logger.Fatal(log.Message(err))
		}
	}
}

func DeployEnv(osArgs []string, factories ...Factory) (env string) {
	if factories != nil && len(factories) != 0 {
		factories[0]()
	} else {
		File(Path("."), Name("config"), Type("toml"))()
	}

	if len(osArgs) == 2 {
		env = osArgs[1]
	} else {
		env = "dev"
	}
	conf = viper.GetViper().Sub(env)

	if conf == nil {
		logger.Fatal("the config of deploy env is nil")
	}
	return
}
