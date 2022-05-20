package config

import (
	"github.com/allentom/harukap/config"
	"os"
)

var DefaultConfigProvider *config.Provider

func InitConfigProvider() error {
	var err error
	customConfigPath := os.Getenv("APP_CONFIG_PATH")
	DefaultConfigProvider, err = config.NewProvider(func(provider *config.Provider) {
		ReadConfig(provider)
	}, customConfigPath)
	return err
}

var Instance Config

type Config struct {
	Port       string
	StaticPath string
	HTMLSource string
	APIProxy   APIProxy
}
type APIProxy struct {
	Enable  bool
	Prefix  string
	Proxy   string
	Rewrite bool
}

func ReadConfig(provider *config.Provider) {
	configer := provider.Manager
	configer.SetDefault("addr", ":8000")
	configer.SetDefault("application", "My Service")
	configer.SetDefault("instance", "main")
	configer.SetDefault("staticpath", "./static")
	configer.SetDefault("htmlsource", "./static/index.html")
	configer.SetDefault("apiproxy.enable", false)
	Instance = Config{
		Port:       configer.GetString("port"),
		StaticPath: configer.GetString("staticpath"),
		APIProxy: APIProxy{
			Enable:  configer.GetBool("apiproxy.enable"),
			Prefix:  configer.GetString("apiproxy.prefix"),
			Proxy:   configer.GetString("apiproxy.proxy"),
			Rewrite: configer.GetBool("apiproxy.rewrite"),
		},
	}
}
