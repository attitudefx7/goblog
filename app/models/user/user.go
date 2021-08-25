package user

import (
	"github.com/attitudefx7/goblog/app/models"
	"github.com/attitudefx7/goblog/pkg/password"
	"github.com/attitudefx7/goblog/pkg/route"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(_password string) bool  {
	return password.CheckHash(_password, u.Password)
}

func (u User) Link() string {
	return route.Name2URL("users.show", "id", u.GetStringID())
}
