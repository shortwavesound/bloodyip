package bloodyip

import "testing"

func TestIsMyIPBloody(t *testing.T) {
	res, err := IsMyIPBloody()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("isBloody = %v\n", res)
}

func TestIsMyIPBloodyWithBelarus(t *testing.T) {
	CheckBelarus = true
	res, err := IsMyIPBloody()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("CheckBelarus = %v, isBloody = %v\n", CheckBelarus, res)
}
