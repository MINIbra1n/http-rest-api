package store

import "github.com/MINIbra1n/http-rest-api/internal/app/model"

// User ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
