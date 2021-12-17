package services

import (
	"go-server-template/models"
	"go-server-template/pkg/database"
	"go-server-template/repositories"
)

type userService struct{}

var UserService = new(userService)

func (s *userService) GetUserInfo(id int64) *models.User {
	return repositories.UserRepository.FindOne(database.DB(), id)
}
