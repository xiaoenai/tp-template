package logic

import (
	tp "github.com/henrylee2cn/teleport"
	micro "github.com/xiaoenai/tp-micro"
	"github.com/xiaoenai/tp-template/accessToken"
	"github.com/xiaoenai/tp-template/account/args"
	"github.com/xiaoenai/tp-template/account/logic/model"
)

// 增加用户
func V1_User_Set(ctx tp.PullCtx, arg *args.SetUserArgsV1) (*args.SetUserResultV1, *tp.Rerror) {
	id, err := model.InsertUser(&model.User{
		Name: arg.Name,
	})
	if err != nil {
		return nil, micro.RerrInternalServerError.SetReason(err.Error())
	}
	// create access token
	token := accessToken.Create(id, "")

	// select user
	user, _, err := model.GetUserByPrimary(id)
	if err != nil {
		return nil, micro.RerrInternalServerError.SetReason(err.Error())
	}
	if user == nil {
		return nil, micro.RerrNotFound
	}

	// update
	user.AccessToken = token.String()
	if err := model.UpdateUserByPrimary(user, []string{"access_token"}); err != nil {
		return nil, micro.RerrInternalServerError.SetReason(err.Error())
	}
	return new(args.SetUserResultV1), nil
}

// 根据ID获取user
func V1_User_GetById(ctx tp.PullCtx, arg *args.GetUserByIdArgsV1) (*args.GetUserByIdResultV1, *tp.Rerror) {
	user, exists, err := model.GetUserByPrimary(arg.Id)
	if err != nil {
		return nil, micro.RerrInternalServerError.SetReason(err.Error())
	}
	if !exists {
		return nil, nil
	}
	return model.ToArgsUser(user), nil
}
