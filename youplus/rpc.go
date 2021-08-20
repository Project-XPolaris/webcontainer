package youplus

import (
	"context"
	"github.com/project-xpolaris/youplustoolkit/youplus/rpc"
	"github.com/spf13/viper"
	"time"
)

var DefaultRPCClient *rpc.YouPlusRPCClient

func LoadYouPlusRPCClient() error {
	DefaultRPCClient = rpc.NewYouPlusRPCClient(viper.GetString("youplus.rpc.url"))
	DefaultRPCClient.KeepAlive = viper.GetBool("youplus.rpc.autoReconnect")
	DefaultRPCClient.MaxRetry = viper.GetInt("youplus.rpc.maxRetry")
	timeoutCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return DefaultRPCClient.Connect(timeoutCtx)
}
