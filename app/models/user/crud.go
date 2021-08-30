package user

import (
	"github.com/attitudefx7/goblog/pkg/logger"
	"github.com/attitudefx7/goblog/pkg/model"
	"github.com/attitudefx7/goblog/pkg/types"
)

func (u *User) Create() (err error) {
	if err = model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(idstr string) (User, error) {
	var user User
	id := types.StringToInt(idstr)

	if err := model.DB.First(&user, id).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.First(&user, "email", email).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func All() ([]User, error) {
	var users []User
	if err := model.DB.Find(&users).Order("created_at desc").Error; err != nil {
		logger.LogError(err)
		return users, err
	}

	return users, nil
}
