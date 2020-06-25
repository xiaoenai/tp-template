package authToken

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/utils"
)

// AUTH_TOKEN_KEY auth token key
const AUTH_TOKEN_KEY = "auth_token_"

var (
	invalidToken = errors.New("无效的auth_token")
)

// Create
func Create(uid int64) *AuthToken {
	return &AuthToken{
		uidInt64: uid,
		token:    strconv.FormatInt(uid, 30),
	}
}

// Secret returns the sign secret.
func (a *AuthToken) Secret() string {
	if a.secret == "" {
		return fmt.Sprintf("%d", time.Now().Unix())
	}
	return a.secret
}

// SetSecret set sign sceret
func (a *AuthToken) SetSecret(secret string) {
	a.secret = secret
}

// Parse 解析auth_token
func Parse(token string) (obj *AuthToken, err error) {
	i, err := strconv.ParseInt(token, 30, 32)
	if err != nil {
		return
	}
	obj = &AuthToken{
		uid:      strconv.FormatInt(i, 10),
		uidInt64: i,
		token:    token,
	}
	return
}

type (
	// AuthToken auth token info
	AuthToken struct {
		uid      string
		uidInt64 int64
		token    string
		secret   string
		info     utils.Args
	}
	// Builder Verifies auth token
	Builder func(query *utils.Args) (*AuthToken, *tp.Rerror)
)

// String returns the auth token string.
func (a *AuthToken) String() string {
	return a.token
}

// SessionId specifies the string as the session ID.
func (a *AuthToken) SessionId() string {
	return a.uid
}

// Uid returns the user id.
func (a *AuthToken) Uid() string {
	return a.uid
}

// UidInt64 returns the user id number.
func (a *AuthToken) UidInt64() int64 {
	return a.uidInt64
}

// AddedQuery the user information will be appended to the URI query part.
func (a *AuthToken) AddedQuery() *utils.Args {
	return &a.info
}
