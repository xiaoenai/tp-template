package accessToken

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/utils"
)

// ACCESS_TOKEN_KEY auth token key
const ACCESS_TOKEN_KEY = "access_token_"

var (
	invalidToken = errors.New("无效的access_token")
)

// Create
func Create(uid int64, deviceId string) *AccessToken {
	return &AccessToken{
		uidInt64: uid,
		token:    strconv.FormatInt(uid, 30),
		// 可使用deviceid与uid一起生成
		deviceId: deviceId,
	}
}

// Secret returns the sign secret.
func (a *AccessToken) Secret() string {
	if a.secret == "" {
		return fmt.Sprintf("%d", time.Now().Unix())
	}
	return a.secret
}

// SetSecret set sign sceret
func (a *AccessToken) SetSecret(secret string) {
	a.secret = secret
}

// Parse 解析auth_token
func Parse(token string) (obj *AccessToken, err error) {
	i, err := strconv.ParseInt(token, 30, 32)
	if err != nil {
		return
	}
	obj = &AccessToken{
		uid:      strconv.FormatInt(i, 10),
		uidInt64: i,
		token:    token,
	}
	return
}

type (
	// AccessToken auth token info
	AccessToken struct {
		uid      string
		uidInt64 int64
		token    string
		secret   string
		deviceId string
		info     utils.Args
	}
	// Builder Verifies auth token
	Builder func(query *utils.Args) (*AccessToken, *tp.Rerror)
)

// String returns the access token string.
func (a *AccessToken) String() string {
	return a.token
}

// SessionId specifies the string as the session ID.
func (a *AccessToken) SessionId() string {
	return a.uid
}

// Uid returns the user id.
func (a *AccessToken) Uid() string {
	return a.uid
}

// UidInt64 returns the user id number.
func (a *AccessToken) UidInt64() int64 {
	return a.uidInt64
}

// AddedQuery the user information will be appended to the URI query part.
func (a *AccessToken) AddedQuery() *utils.Args {
	return &a.info
}

// DeviceId returns the user id number.
func (a *AccessToken) DeviceId() string {
	return a.deviceId
}
