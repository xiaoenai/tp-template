package main

import (
	"time"

	"github.com/henrylee2cn/cfgo"
	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/model/etcd"
	"github.com/xiaoenai/tp-micro/model/mysql"
	"github.com/xiaoenai/tp-micro/model/redis"

	"github.com/xiaoenai/tp-template/account/logic/model"
)

type config struct {
	Srv         micro.SrvConfig `yaml:"srv"`
	Etcd        etcd.EasyConfig `yaml:"etcd"`
	Mysql       mysql.Config    `yaml:"mysql"`
	Redis       redis.Config    `yaml:"redis"`
	CacheExpire time.Duration   `yaml:"cache_expire"`
	LogLevel    string          `yaml:"log_level"`
}

func (c *config) Reload(bind cfgo.BindFunc) error {
	err := bind()
	if err != nil {
		return err
	}
	if c.CacheExpire == 0 {
		c.CacheExpire = time.Hour * 24
	}
	if len(c.LogLevel) == 0 {
		c.LogLevel = "TRACE"
	}
	tp.SetLoggerLevel(c.LogLevel)
	var (
		mysqlConfig *mysql.Config
		redisConfig = &c.Redis
	)
	if len(c.Mysql.Host) > 0 {
		mysqlConfig = &c.Mysql
	}
	err = model.Init(mysqlConfig, nil, redisConfig, c.CacheExpire)
	if err != nil {
		tp.Errorf("%v", err)
	}
	return nil
}

var cfg = &config{
	Srv: micro.SrvConfig{
		ListenAddress:     ":0",
		EnableHeartbeat:   true,
		PrintDetail:       true,
		CountTime:         true,
		SlowCometDuration: time.Millisecond * 500,
	},
	Etcd: etcd.EasyConfig{
		Endpoints: []string{"http://127.0.0.1:2379"},
	},
	Redis:       *redis.NewConfig(),
	CacheExpire: time.Hour * 24,
	LogLevel:    "TRACE",
}

func init() {
	goutil.WritePidFile()
	cfgo.MustReg("account", cfg)
}
