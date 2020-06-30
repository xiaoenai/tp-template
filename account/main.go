package main

import (
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/discovery"
	"github.com/xiaoenai/tp-template/account/api"
	"github.com/xiaoenai/tp-template/plugin"
)

func main() {
	srv := micro.NewServer(
		cfg.Srv,
		discovery.ServicePlugin(cfg.Srv.InnerIpPort(), cfg.Etcd),
		plugin.NewInnerAuth(true),
	)
	api.Route("/account", srv.Router())
	srv.ListenAndServe()
}
