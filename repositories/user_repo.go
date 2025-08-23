package repositories

import (
	"go-fiber-api/config"
	"go-fiber-api/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := config.DB.Unscoped().First(&user, id)
	return user, result.Error
}

func CreateUser(user models.User) (models.User, error) {
	result := config.DB.Create(&user)
	return user, result.Error
}

func UpdateUser(id uint, updatedUser models.User) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	user.Name = updatedUser.Name
	user.Username = updatedUser.Username
	user.ProvinceCode = updatedUser.ProvinceCode
	user.CityCode = updatedUser.CityCode
	user.DistrictCode = updatedUser.DistrictCode
	user.RoleID = updatedUser.RoleID
	config.DB.Save(&user)
	return user, nil
}

func DeleteUser(id uint) error {
	result := config.DB.Delete(&models.User{}, id)
	return result.Error
}
