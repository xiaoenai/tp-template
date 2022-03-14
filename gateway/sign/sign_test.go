package sign

import (
	"testing"

	"github.com/henrylee2cn/teleport/utils"
)

func TestSign(t *testing.T) {
	secret := "1593065127"
	args := utils.AcquireArgs()
	args.Set("a", "1")
	args.Set("b", "2")
	body := []byte("body123")

	// 签名
	sign := Sign(secret, args, body)
	args.Set(SIGN_KEY, sign)
	t.Logf("signed: %s", args.String())

	// 校验签名
	if !VerifySign(secret, args, body) {
		t.Fail()
	}
}
