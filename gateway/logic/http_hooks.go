package logic

import (
	"bytes"

	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/socket"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-micro/gateway/types"
)

type (
	httpHooks struct{}
)

// OnRequest is called when the client requests.
func (*httpHooks) OnRequest(params types.RequestArgs, body []byte, fn types.AuthFunc) (types.AccessToken, []socket.PacketSetting, *tp.Rerror) {
	var (
		args       = params.QueryArgs()
		hasIllegal bool
	)
	// 防止伪造内部参数
	args.VisitAll(func(key, _ []byte) {
		if bytes.HasPrefix(key, []byte{'_'}) {
			hasIllegal = true
		}
	})
	if hasIllegal {
		return nil, nil, micro.RerrInvalidParameter
	}

	// 校验签名
	t, rerr := autoVerifySign(params, body, fn)
	if rerr != nil || t == nil {
		return nil, nil, rerr
	}
	return t, nil, nil
}
