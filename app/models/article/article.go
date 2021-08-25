package article

import (
	"github.com/attitudefx7/goblog/app/models"
	"github.com/attitudefx7/goblog/app/models/user"
	"github.com/attitudefx7/goblog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(50);" valid:"title"`
	Body  string `gorm:"type:varchar(500);" valid:"body"`

	UserId uint64 `gorm:"not null;index"`
	User user.User

	CategoryID uint64 `gorm:"not null; default:4;index"`
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
