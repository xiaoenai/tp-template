package accessToken

import (
	"testing"
)

func TestAccessToken(t *testing.T) {
	// gen
	token := Create(100, "")
	t.Logf("Token-> %#v", token)

	// parse
	parseToken, err := Parse(token.String())
	if err != nil {
		t.Fatalf("Parse err-> %v", err)
	}
	t.Logf("Parse Token-> %#v", parseToken)
}
