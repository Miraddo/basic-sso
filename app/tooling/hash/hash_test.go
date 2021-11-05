package hash_test

import (
	"testing"

	"github.com/miraddo/basic-sso/app/tooling/hash"
)

func TestToken(t *testing.T) {
	t.Log("Test Token Generator")
	{
		token, err := hash.GenToken()

		switch {
		case err != nil:
			t.Errorf("while generate token got an error %v", err)

		case len(token) != 120:
			t.Errorf("the length should be 60 but we got this [%v] with length [%d]", token, len(token))

		default:
			t.Log("Everything is fine")
		}
	}
}

func TestGenPass(t *testing.T) {
	t.Log("Test Password Generator")
	{
		data := []struct {
			password string
			expected string
		}{
			{"1234", "81dc9bdb52d04dc20036dbd8313ed055"},
			{"12345", "827ccb0eea8a706c4c34a16891f84e7b"},
			{"password", "5f4dcc3b5aa765d61d8327deb882cf99"},
			{"pass", "1a1dc91c907325c69271ddf0c944bc72"},
		}

		for _, x := range data {
			hash := hash.GenPass(x.password)

			if hash != x.expected {
				t.Errorf("Expected [%v], but got [%v]", x.expected, hash)
			}

			t.Log("Everything is fine")
		}

	}
}
