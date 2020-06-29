package gateway

import (
	"time"

	"github.com/henrylee2cn/cfgo"
	tp "github.com/henrylee2cn/teleport"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/gateway"
	short "github.com/xiaoenai/tp-micro/gateway/logic/http"
	"github.com/xiaoenai/tp-micro/model/etcd"
	"github.com/xiaoenai/tp-micro/model/mysql"
	"github.com/xiaoenai/tp-micro/model/redis"
)

type Config struct {
	Gw               gateway.Config  `yaml:"gw"`
	GraySocketClient micro.CliConfig `yaml:"gray_socket_client"`
	GrayEtcd         etcd.EasyConfig `yaml:"gray_etcd"`
	Redis            redis.Config    `yaml:"redis"`
	Mysql            mysql.Config    `yaml:"mysql"`
	// micro hosts save key
	HostsPrefix string `yaml:"hosts_prefix"`
	// 长连接redis订阅上下线推送
	Namespace string `yaml:"namespace"`
	LogLevel  string `yaml:"log_level"`
}

func (c *Config) Reload(bind cfgo.BindFunc) error {
	err := bind()
	if err == nil {
		c.Gw.OuterHttpServer.OuterIpPort()
	}
	if len(c.LogLevel) == 0 {
		c.LogLevel = "TRACE"
	}
	tp.SetLoggerLevel(c.LogLevel)

	c.Gw.OuterHttpServer.PrintDetail = c.Gw.OuterSocketServer.PrintDetail
	c.Gw.OuterHttpServer.CountTime = c.Gw.OuterSocketServer.CountTime
	c.Gw.OuterHttpServer.SlowCometDuration = c.Gw.OuterSocketServer.SlowCometDuration
	return nil
}

var cfg = Config{
	Gw: gateway.Config{
		EnableHttp:   true,
		EnableSocket: true,
		OuterHttpServer: short.HttpSrvConfig{
			ListenAddress: "0.0.0.0:5000",
			AllowCross:    true,
		},
		OuterSocketServer: micro.SrvConfig{
			ListenAddress:     "0.0.0.0:5020",
			EnableHeartbeat:   true,
			PrintDetail:       true,
			CountTime:         true,
			SlowCometDuration: time.Millisecond * 500,
		},
		InnerSocketServer: micro.SrvConfig{
			ListenAddress:     "0.0.0.0:5030",
			EnableHeartbeat:   true,
			PrintDetail:       true,
			CountTime:         true,
			SlowCometDuration: time.Millisecond * 500,
		},
		InnerSocketClient: micro.CliConfig{
			Failover:        3,
			HeartbeatSecond: 60,
		},
		Etcd: etcd.EasyConfig{
			Endpoints: []string{"http://127.0.0.1:2379"},
		},
	},
	GraySocketClient: micro.CliConfig{
		Failover:        3,
		HeartbeatSecond: 60,
	},
	GrayEtcd: etcd.EasyConfig{
		Endpoints: []string{"http://127.0.0.1:2379"},
	},
	Redis:       *redis.NewConfig(),
	Mysql:       *mysql.NewConfig(),
	HostsPrefix: "MICRO-GW_HOSTS",
	LogLevel:    "TRACE",
}

func init() {
	cfgo.MustReg("gateway", &cfg)
}
