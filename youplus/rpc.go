package youplus

import (
	"github.com/project-xpolaris/youplustoolkit/youplus/rpc"
	"github.com/spf13/viper"
)

var DefaultRPCClient *rpc.YouPlusRPCClient

func LoadYouPlusRPCClient() error {
	DefaultRPCClient = rpc.NewYouPlusRPCClient(viper.GetString("youplus.rpc.url"))
	return DefaultRPCClient.Init()
}
