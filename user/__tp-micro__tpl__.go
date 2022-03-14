// Command user is the tp-micro service project.
// The framework reference: https://github.com/xiaoenai/tp-micro
package __TPL__

import "github.com/henrylee2cn/teleport/codec"

// __API_PULL__ register PULL router
type __API_PULL__ interface {
	V1_User
}

// __API_PUSH__ register PUSH router:
//  /stat
type __API_PUSH__ interface {
}

// V1_User
type V1_User interface {
	Add(*AddUserArgsV1) *AddUserResultV1
	// 获取用户
	GetById(*GetUserByIdArgsV1) *GetUserByIdResultV1
	Get(*GetUserArgsV1) *GetUserResultV1
}

type (
	AddUserArgsV1 struct {
		Name string
	}
	AddUserResultV1 = codec.PbEmpty

	GetUserByIdArgsV1 struct {
		Id int64 `param:"<query> <nonzero> <rerr:400:id不能为空>"`
	}
	GetUserByIdResultV1 struct {
		Name        string
		AccessToken string
	}

	GetUserArgsV1   = codec.PbEmpty
	GetUserResultV1 struct {
		Name string
	}
)

// __MYSQL_MODEL__ create mysql model
type __MYSQL_MODEL__ struct {
}

// __MONGO_MODEL__ create mongodb model
type __MONGO_MODEL__ struct {
}
