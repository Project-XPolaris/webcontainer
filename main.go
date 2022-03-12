package main

import (
	"github.com/allentom/harukap"
	"github.com/allentom/harukap/cli"
	"github.com/sirupsen/logrus"
	"webcontainer/config"
	"webcontainer/plugin"
	"webcontainer/youlog"
	"webcontainer/youplus"
)

func main() {
	err := config.InitConfigProvider()
	if err != nil {
		logrus.Fatal(err)
	}
	err = youlog.DefaultYouLogPlugin.OnInit(config.DefaultConfigProvider)
	if err != nil {
		logrus.Fatal(err)
	}
	appEngine := harukap.NewHarukaAppEngine()
	appEngine.ConfigProvider = config.DefaultConfigProvider
	appEngine.LoggerPlugin = youlog.DefaultYouLogPlugin
	appEngine.UsePlugin(youplus.DefaultYouPlusPlugin)
	appEngine.UsePlugin(&plugin.LauncherPlugin{})
	if err != nil {
		logrus.Fatal(err)
	}
	appWrap, err := cli.NewWrapper(appEngine)
	if err != nil {
		logrus.Fatal(err)
	}
	appWrap.RunApp()
}
