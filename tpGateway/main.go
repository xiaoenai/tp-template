package main

import (
	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/utils"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/discovery"
	"github.com/xiaoenai/tp-template/accessToken"
	accountArgs "github.com/xiaoenai/tp-template/account/args"
	accountSdk "github.com/xiaoenai/tp-template/account/sdk"
	"github.com/xiaoenai/tp-template/gateway"
)

func main() {
	gateway.Main()
}

func init() {
	// 初始化account sdk
	client := micro.NewClient(
		gateway.GetConfig().Gw.InnerSocketClient,
		discovery.NewLinker(gateway.GetConfig().Gw.Etcd),
	)
	accountSdk.InitWithClient(client)

	var (
		requiredQueryParams = []string{"system", "app_ver"}
	)

	gateway.RegBuilder("tp", func(query *utils.Args) (*accessToken.AccessToken, *tp.Rerror) {
		// 检查必需的参数
		for _, p := range requiredQueryParams {
			if !query.Has(p) {
				return nil, micro.RerrInvalidParameter
			}
		}

		token := query.Peek(accessToken.ACCESS_TOKEN_KEY)
		if len(token) == 0 {
			return nil, nil
		}

		obj, err := accessToken.Parse(goutil.BytesToString(token))
		if err != nil {
			tp.Errorf("parse access_token error, token-> %s, err-> %s", token, err.Error())
			return nil, micro.RerrInvalidParameter.SetMessage("access token 错误")
		}

		// 检查用户正确性
		_user, rerr := accountSdk.V1_User_GetById(&accountArgs.GetUserByIdArgsV1{
			Id: obj.UidInt64(),
		})

		if rerr != nil {
			return nil, rerr
		}

		// 签名参数设置
		obj.SetSecret("secret")

		if goutil.BytesToString(token) != _user.AccessToken {
			return nil, micro.RerrInvalidParameter.SetMessage("access token 非法")
		}
		return obj, nil
	})
}
