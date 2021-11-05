package database_test

import (
	"testing"

	"github.com/miraddo/basic-sso/app/tooling/database"
)

func TestDatabase(t *testing.T) {
	t.Log("Test Database Insert")
	{
		data := []struct {
			username string
			password string
			expected bool
		}{
			{"user1", "pass1", true},
			{"user2", "pass2", true},
		}

		for _, x := range data {
			user := database.User{
				Username: x.username,
				Password: x.password,
			}

			if result, err := user.Insert(); err != nil {
				t.Errorf("expected [%v] , but got [%v]", x.expected, result)
			}

			t.Log("Everything looks good")
		}
	}
}
