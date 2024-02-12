package repository

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type inMemoryUserRepository struct {
	users map[string]entity.User
}

func InMemoryUserRepository() UserRepository {
	users := make(map[string]entity.User, 0)
	return &inMemoryUserRepository{users: users}
}

// Save implements UserRepository.
func (repo *inMemoryUserRepository) Save(userRequest *entity.User) error {
	_, exist := repo.users[userRequest.Username]
	if exist {
		return custerror.OccupiedUsernameError(userRequest.Username)
	}
	repo.users[userRequest.Username] = *userRequest

	return nil
}

func (repo *inMemoryUserRepository) FindByUsername(username string) (*entity.User, error) {
	user, ok := repo.users[username]
	if !ok {
		return nil, custerror.UserNotFoundError(username)
	}
	return &user, nil
}
