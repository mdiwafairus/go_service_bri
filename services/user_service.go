package services

import (
	"go-fiber-api/config"
	"go-fiber-api/models"
	"go-fiber-api/repositories"
)

func GetUsers(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * limit

	if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	result := config.DB.
		Limit(limit).
		Offset(offset).
		Find(&users)

	return users, total, result.Error
}

func GetUser(id uint) (models.User, error) {
	return repositories.GetUserByID(id)
}

func CreateUser(user models.User) (models.User, error) {
	return repositories.CreateUser(user)
}

func UpdateUser(id uint, user models.User) (models.User, error) {
	return repositories.UpdateUser(id, user)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}
