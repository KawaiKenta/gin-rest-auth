package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"size:255" json:"name,omitempty"`
	Email      string `gorm:"size:255;not null;unique" json:"email,omitempty"`
	Password   string `gorm:"size:255;not null" json:"password,omitempty"`
	IsVerified bool   `gorm:"not null" json:"is_verified"`
}

func GetUserByEmail(email string) (*User, error) {
	var user *User
	if err := DB.Where("email=?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateNewUser(user *User) error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(id int) (*User, error) {
	var user *User
	if err := DB.Where("ID=?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeleteUser(id int) error {
	result := DB.Where("id=?", id).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("record not found")
	}
	return nil
}

func UpdateUser(user *User) error {
	result := DB.Model(&user).Updates(map[string]interface{}{
		"name":     user.Name,
		"password": user.Password,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("record not found")
	}
	return nil
}
