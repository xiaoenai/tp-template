package plugin

import (
	"strings"

	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/socket"

	"github.com/xiaoenai/tp-template/gateway/logic"
)

const (
	prefixMark      = logic.PASS_KEY + "="
	middleMark      = "&" + logic.PASS_KEY + "="
	prefixInnerMark = logic.PASS_KEY + "=56FFFDF4C2EBACFD"
	middleInnerMark = "&" + logic.PASS_KEY + "=56FFFDF4C2EBACFD"
)

// AddInnerAuth 请求时添加内部访问权限
func AddInnerAuth(settings ...socket.PacketSetting) socket.PacketSetting {
	innerFn := tp.WithQuery(logic.PASS_KEY, "56FFFDF4C2EBACFD")
	return func(p *socket.Packet) {
		innerFn(p)
		for _, fn := range settings {
			fn(p)
		}
	}
}

type (
	authFilter struct{ onlyInner bool }
)

// NewInnerAuth 创建访问权限插件
func NewInnerAuth(onlyInner ...bool) tp.Plugin {
	p := new(authFilter)
	if len(onlyInner) > 0 {
		p.onlyInner = onlyInner[0]
	}
	return p
}

var (
	_ tp.PreReadPullBodyPlugin = (*authFilter)(nil)
	_ tp.PreReadPushBodyPlugin = (*authFilter)(nil)
)

// Name
func (a *authFilter) Name() string {
	if a.onlyInner {
		return "inner_auth"
	}
	return "gw_auth"
}

// PreReadPullBody
func (a *authFilter) PreReadPullBody(ctx tp.ReadCtx) *tp.Rerror {
	queryString := ctx.UriObject().RawQuery
	if a.onlyInner {
		if strings.HasPrefix(queryString, prefixInnerMark) ||
			strings.Contains(queryString, middleInnerMark) {
			return nil
		}
	} else {
		if strings.HasPrefix(queryString, prefixMark) ||
			strings.Contains(queryString, middleMark) {
			return nil
		}
	}
	return tp.NewRerror(tp.CodeUnauthorized, "内部认证失败", "")
}

// PreReadPushBody
func (a *authFilter) PreReadPushBody(ctx tp.ReadCtx) *tp.Rerror {
	return a.PreReadPullBody(ctx)
}
