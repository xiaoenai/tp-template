// Code generated by 'micro gen' command.
// DO NOT EDIT!

package args

import (
	"github.com/henrylee2cn/teleport/codec"
)

// EmptyStruct alias of type struct {}
type EmptyStruct = codec.PbEmpty

// SetUserResultV1 alias of type codec.PbEmpty
type SetUserResultV1 = codec.PbEmpty

// GetUserByIdResultV1 alias of type User
type GetUserByIdResultV1 = User

// SetUserArgsV1 comment...
type SetUserArgsV1 struct {
	Name string `json:"name"`
}

// GetUserByIdArgsV1 comment...
type GetUserByIdArgsV1 struct {
	Id int64 `json:"id"`
}

// User user info
type User struct {
	Id          int64  `key:"pri" json:"id"`
	Name        string `key:"uni" json:"name"`
	AccessToken string `json:"access_token"`
	UpdatedAt   int64  `json:"updated_at"`
	CreatedAt   int64  `json:"created_at"`
	DeletedTs   int64  `json:"deleted_ts"`
}
