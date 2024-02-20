package repository

import (
	"github.com/lkcsi/bookstore/custerror"
	"github.com/lkcsi/bookstore/entity"
)

type inMemoryUserRepository struct {
	users map[string]entity.User
}

func ImUserRepository() UserRepository {
	users := make(map[string]entity.User, 0)
	user := entity.User{Username: "test", Password: "$2a$10$KJKTiTcOhHjVIVH74u8pCOv18tzOs4Fd8bd8Dl7mZlJy/q2Tj2Vjq"}
	users[user.Username] = user
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
