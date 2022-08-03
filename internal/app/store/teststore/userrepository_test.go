package teststore_test

import (
	"github.com/Dennikoff/UserTagApi/internal/app/model"
	"github.com/Dennikoff/UserTagApi/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	st := teststore.New()
	testcases := []struct {
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
			name: "Duplicate key",
			user: func() *model.User {
				return model.TestUser()
			},
			ok: false,
		},
		{
			name: "invalid pass",
			user: func() *model.User {
				user := model.TestUser()
				user.Password = ""
				return user
			},
			ok: false,
		},
		{
			name: "invalid email 1",
			user: func() *model.User {
				user := model.TestUser()
				user.Email = "test@email"
				return user
			},
			ok: false,
		},
		{
			name: "invalid email 2",
			user: func() *model.User {
				user := model.TestUser()
				user.Email = "testemail.com"
				return user
			},
			ok: false,
		},
	}
	for _, cs := range testcases {
		t.Run(cs.name, func(t *testing.T) {
			if cs.ok {
				assert.NoError(t, st.User().Create(cs.user()))
			} else {
				assert.Error(t, st.User().Create(cs.user()))
			}
		})
	}

}