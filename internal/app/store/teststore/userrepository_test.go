package teststore_test

import (
	"testing"

	"github.com/MINIbra1n/http-rest-api/internal/app/model"
	"github.com/MINIbra1n/http-rest-api/internal/app/store"
	"github.com/MINIbra1n/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create((u)))
	assert.NotNil(t, u)

}

func TestUserRepositoryFindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(model.TestUser(t))
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
func TestUserRepositoryFind(t *testing.T) {

	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)

}
