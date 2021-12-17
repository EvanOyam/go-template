package repositories

import (
	"go-server-template/models"

	"gorm.io/gorm"
)

type userRepository struct{}

var UserRepository = new(userRepository)

func (r *userRepository) FindOne(db *gorm.DB, id int64) *models.User {
	user := &models.User{}
	if err := db.First(user, "id = ?", id).Error; err != nil {
		return nil
	}
	return user
}
