package sqlstore_test

import (
	"github.com/Dennikoff/UserTagApi/internal/app/model"
	"github.com/Dennikoff/UserTagApi/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, del := sqlstore.TestDB(t, DatabaseURL)

	st := sqlstore.New(db)

	defer del("users")

	user := model.TestUser()

	assert.NoError(t, st.User().Create(user))
	assert.Error(t, st.User().Create(user))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, del := sqlstore.TestDB(t, DatabaseURL)

	st := sqlstore.New(db)

	defer del("users")

	user := model.TestUser()

	assert.NoError(t, st.User().Create(user))
	u, _ := st.User().FindByEmail(user.Email)

	assert.NotNil(t, u)
	u, err := st.User().FindByEmail("not@in.db")
	assert.Error(t, err)
	assert.Nil(t, u)
}
