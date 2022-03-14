package sdk

import (
	"strconv"

	tp "github.com/henrylee2cn/teleport"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-template/account/args"
)

var (
	_ tp.PostReadPullBodyPlugin = (*AccountPlugin)(nil)
)

type (
	AccountPlugin struct{}
)

func NewAccountPlugin() *AccountPlugin {
	return nil
}

func (*AccountPlugin) Name() string {
	return "AccountPlugin"
}

type (
	userInfoKey string
)

const (
	QUERY_UID                    = "_uid"
	CTX_DATAKEY_USER userInfoKey = "__ACCOUNT__"
)

// PostReadPullBody
func (u *AccountPlugin) PostReadPullBody(ctx tp.ReadCtx) *tp.Rerror {
	uidStr := ctx.Query().Get(QUERY_UID)
	if len(uidStr) == 0 {
		return micro.RerrInvalidParameter.SetMessage("未找到_uid参数")
	}

	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return micro.RerrInvalidParameter.SetMessage("从_uid参数中解析uid失败 " + err.Error())
	}
	// 获取用户信息
	user, rerr := V1_User_GetById(
		&args.GetUserByIdArgsV1{
			Id: uid,
		},
	)
	if rerr != nil {
		return rerr
	}
	ctx.Swap().Store(CTX_DATAKEY_USER, user)
	return nil
}

// PostReadPushBody
func (u *AccountPlugin) PostReadPushBody(ctx tp.ReadCtx) *tp.Rerror {
	return u.PostReadPullBody(ctx)
}
