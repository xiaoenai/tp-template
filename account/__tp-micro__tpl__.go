// Command account is the tp-micro service project.
// The framework reference: https://github.com/xiaoenai/tp-micro
package __TPL__

import "github.com/henrylee2cn/teleport/codec"

// __API_PULL__ register PULL router
type __API_PULL__ interface {
	V1_User
}

// V1_User
type V1_User interface {
	// 增加用户
	Set(*SetUserArgsV1) *SetUserResultV1
	// 根据ID获取user
	GetById(*GetUserByIdArgsV1) *GetUserByIdResultV1
}

type (
	SetUserArgsV1 struct {
		Name string
	}
	SetUserResultV1 = codec.PbEmpty

	GetUserByIdArgsV1 struct {
		Id int64
	}
	GetUserByIdResultV1 = User
)

// __API_PUSH__ register PUSH router:
type __API_PUSH__ interface {
}

// __MYSQL_MODEL__ create mysql model
type __MYSQL_MODEL__ struct {
	User
}

// __MONGO_MODEL__ create mongodb model
type __MONGO_MODEL__ struct {
}

// User user info
type User struct {
	Id          int64  `key:"pri"`
	Name        string `key:"uni"`
	AccessToken string
}
