package sign

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strings"

	"github.com/henrylee2cn/goutil"
	tp "github.com/henrylee2cn/teleport"
	"github.com/henrylee2cn/teleport/utils"
)

const (
	// 固定字段名
	SIGN_KEY = "sig_"
)

// VerifySign 验证签名
func VerifySign(secret string, args *utils.Args, body []byte) bool {
	// 获取请求签名参数
	oldSign := goutil.BytesToString(args.Peek(SIGN_KEY))

	// 计算签名
	newSign := Sign(secret, args, body)
	matched := oldSign == newSign
	if !matched {
		tp.Debugf("verification failed: old=%s, but new=%s", oldSign, newSign)
	}
	return matched
}

// Sign 计算签名
func Sign(secret string, args *utils.Args, body []byte) string {
	defer args.SetBytesV(SIGN_KEY, args.Peek(SIGN_KEY))
	args.Del(SIGN_KEY)

	var val = make(url.Values, 20)
	args.VisitAll(func(k, v []byte) {
		val.Add(goutil.BytesToString(k), goutil.BytesToString(v))
	})
	baseString := secret + "&" + strings.Replace(val.Encode(), "+", "%20", -1)
	hash := md5.Sum(append(goutil.StringToBytes(baseString), body...))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}
