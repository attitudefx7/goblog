package category

import (
	"github.com/attitudefx7/goblog/pkg/logger"
	"github.com/attitudefx7/goblog/pkg/model"
	"github.com/attitudefx7/goblog/pkg/types"
)

func (c *Category) Create() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToInt(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		logger.LogError(err)
		return category, err
	}

	return category, nil
}

func All() ([]Category, error) {
	var categories []Category

	if err := model.DB.Find(&categories).Error; err != nil {
		logger.LogError(err)

		return categories, err
	}

	return categories, nil
}
