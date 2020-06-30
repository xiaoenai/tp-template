package main

import (
	"time"

	accountSdk "github.com/xiaoenai/tp-template/account/sdk"

	"github.com/henrylee2cn/cfgo"
	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/model/etcd"
)

type config struct {
	Srv         micro.SrvConfig `yaml:"srv"`
	Etcd        etcd.EasyConfig `yaml:"etcd"`
	Cli         micro.CliConfig `yaml:"cli"`
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

	// init account sdk
	accountSdk.Init(c.Cli, c.Etcd)
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
	CacheExpire: time.Hour * 24,
	LogLevel:    "TRACE",
}

func init() {
	goutil.WritePidFile()
	cfgo.MustReg("user", cfg)
}
