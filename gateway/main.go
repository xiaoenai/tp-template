package gateway

import (
	"strings"

	tp "github.com/henrylee2cn/teleport"
	"github.com/xiaoenai/tp-micro/gateway"
	"github.com/xiaoenai/tp-micro/gateway/helper/agent"
	"github.com/xiaoenai/tp-micro/gateway/helper/gray"
	"github.com/xiaoenai/tp-micro/model/redis"
	"github.com/xiaoenai/tp-template/accessToken"
	"github.com/xiaoenai/tp-template/gateway/logic"
)

// GetConfig returns config
func GetConfig() Config {
	return cfg
}

// Main start gateway
func Main() {
	biz := logic.NewBusiness()

	// new redis client
	redisClient, err := redis.NewClient(&cfg.Redis)
	if err != nil {
		tp.Fatalf("%v", err)
	}

	// socket agent
	agent.Init(redisClient, redisClient, cfg.Namespace)
	biz.SocketHooks = agent.GetSocketHooks()

	// gray
	_, err = gray.SetGray(biz, cfg.GraySocketClient, cfg.GrayEtcd, cfg.Mysql, cfg.Redis, nil)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1142: CREATE command denied to user") {
			tp.Warnf("%v", err)
		} else {
			tp.Fatalf("%v", err)
		}
	}

	// set host key name space
	gateway.SetHostsNamespace(cfg.HostsPrefix)

	// run micro gateway
	go gateway.Run(cfg.Gw, biz, nil)
	select {}
}

// RegBuilder 注册构建AccessToken的函数
// 注：若注册AccessToken的构建函数为nil，或重复注册，均会发生panic
func RegBuilder(authName string, fn accessToken.Builder) {
	logic.RegBuilder(authName, fn)
}
