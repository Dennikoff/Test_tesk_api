package model_test

import (
	"github.com/Dennikoff/UserTagApi/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	users := []struct {
		name string
		user func() *model.User
		ok   bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser()
			},
			ok: true,
		},
		{
			name: "invalid_pass",
			user: func() *model.User {
				user := model.TestUser()
				user.Password = "123"
				return user
			},
			ok: false,
		},
		{
			name: "invalid_email",
			user: func() *model.User {
				user := model.TestUser()
				user.Email = "invalid email.com"
				return user
			},
			ok: false,
		},
		{
			name: "empty nickname",
			user: func() *model.User {
				user := model.TestUser()
				user.Password = ""
				return user
			},
			ok: false,
		},
	}
	for _, us := range users {
		t.Run(us.name, func(t *testing.T) {
			if us.ok {
				assert.NoError(t, us.user().Validate())
			} else {
				assert.Error(t, us.user().Validate())
			}
		})
	}
}
