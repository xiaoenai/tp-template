package logic

import (
	"unsafe"

	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/utils"
	"github.com/xiaoenai/tp-micro/gateway/types"
	"github.com/xiaoenai/tp-template/accessToken"
	"github.com/xiaoenai/tp-template/gateway/sign"
)

// autoVerifySign
func autoVerifySign(params types.RequestArgs, body []byte, fn types.AuthFunc) (*accessToken.AccessToken, *tp.Rerror) {
	args := params.QueryArgs()

	// 获取 accessToken info
	accessTokenIface, rerr := fn(goutil.BytesToString(args.QueryString()))
	if rerr != nil {
		// 允许不进行授权
		if rerr.Code == tp.CodeUnauthorized {
			return nil, nil
		}
		return nil, rerr
	}

	accessToken := accessTokenIface.(*accessToken.AccessToken)

	// 获取请求签名参数
	oldSign := goutil.BytesToString(args.Peek(SIGN_KEY))

	if len(oldSign) == 0 {
		return accessToken, nil
	}

	// 计算签名
	newSign := sign.Sign(accessToken.Secret(), (*utils.Args)(unsafe.Pointer(args)), body)
	matched := oldSign == newSign
	if !matched {
		return nil, tp.NewRerror(tp.CodeUnauthorized, "无效的签名", "")
	}
	return accessToken, nil
}
