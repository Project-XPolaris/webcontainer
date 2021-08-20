package youplus

import (
	"context"
	"fmt"
	entry "github.com/project-xpolaris/youplustoolkit/youplus/entity"
	"github.com/spf13/viper"
)

var DefaultEntry *entry.EntityClient

type AppExport struct {
	Addrs []string `json:"addrs"`
}

func InitEntity() error {
	DefaultEntry = entry.NewEntityClient(
		viper.GetString("youplus.entity.name"),
		viper.GetInt64("youplus.entity.version"),
		&entry.EntityExport{}, DefaultRPCClient,
	)
	DefaultEntry.HeartbeatRate = viper.GetInt64("youplus.entity.heartbeatRate")

	err := DefaultEntry.Register()
	if err != nil {
		return err
	}
	addrs, err := GetHostIpList()
	urls := make([]string, 0)
	for _, addr := range addrs {
		urls = append(urls, fmt.Sprintf("http://%s:%d", addr, viper.GetInt64("port")))
	}
	if err != nil {
		return err
	}
	err = DefaultEntry.UpdateExport(entry.EntityExport{Urls: urls, Extra: map[string]interface{}{}})
	if err != nil {
		return err
	}

	err = DefaultEntry.StartHeartbeat(context.Background())
	if err != nil {
		return err
	}
	return nil
}
