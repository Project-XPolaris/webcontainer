package application

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/webcontainer/config"
	"github.com/rs/cors"
)

func GetHttpService() (*haruka.Engine, error) {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	if config.Instance.APIProxy.Enable {
		e.Router.HandlerRouter.PathPrefix(config.Instance.APIProxy.Prefix).HandlerFunc(apiProxyHandler)
	}
	e.Router.HandlerRouter.PathPrefix("/").Handler(spaHandler{
		staticPath: config.Instance.StaticPath,
		indexPath:  config.Instance.HTMLSource,
	})
	return e, nil
}
