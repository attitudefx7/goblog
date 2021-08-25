package category

import (
	"github.com/attitudefx7/goblog/app/models"
	"github.com/attitudefx7/goblog/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(100);not null;" valid:"name"`
}

func (c Category) Link() string {
	return route.Name2URL("categories.show", "id", c.GetStringID())
}
