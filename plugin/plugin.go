package plugin

import (
	"fmt"
	"github.com/allentom/harukap"
	"net/http"
	"webcontainer/config"
	"webcontainer/filesystem"
)

type LauncherPlugin struct {
}

func (p *LauncherPlugin) OnInit(e *harukap.HarukaAppEngine) error {
	portString := fmt.Sprintf(":%s", config.Instance.Port)
	fs := http.FileServer(&filesystem.FileSystem{Root: http.Dir("static")})
	http.Handle("/", fs)
	return http.ListenAndServe(portString, nil)
}
