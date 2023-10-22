package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique;not null"`
	FirstName     string `gorm:"not null"`
	LastName      string `gorm:"not null"`
	PostalAddress string `gorm:"not null"`
	Email         string `gorm:"unique;not null"`
	PasswordHash  string `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func SoftDeleteUser(db *gorm.DB, id uint) error {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return errors.New("failed to soft delete user")
	}
	return nil
}

func UnDeleteUser(db *gorm.DB, id uint) error {
	user := &User{}
	user.ID = id
	if err := db.Model(user).Update("deleted_at", nil).Error; err != nil {
		return errors.New("failed to undelete user")
	}
	return nil
}

// Your existing functions:

func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func CreateUser(db *gorm.DB, user *User) error {
	if err := db.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func UpdateUser(db *gorm.DB, user *User) error {
	if err := db.Save(user).Error; err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
