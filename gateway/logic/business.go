package logic

import (
	"github.com/xiaoenai/tp-micro/gateway/types"
)

// NewBusiness 创建业务
func NewBusiness() *types.Business {
	b := types.DefaultBusiness()
	b.AuthFunc = authFunc
	b.HttpHooks = new(httpHooks)
	return b
}
