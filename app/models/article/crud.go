package article

import (
	"github.com/attitudefx7/goblog/pkg/logger"
	"github.com/attitudefx7/goblog/pkg/model"
	"github.com/attitudefx7/goblog/pkg/pagination"
	"github.com/attitudefx7/goblog/pkg/route"
	"github.com/attitudefx7/goblog/pkg/types"
	"net/http"
)

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func GetAll(r *http.Request, prepage int) ([]Article, pagination.ViewData, error) {
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), prepage)

	viewData := _pager.Paging()

	var articles []Article

	err := _pager.Results(&articles)

	return articles, viewData, err
}

func GetByUserID(uid string) ([]Article, error) {
	var articles []Article

	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func GetByCategoryID(cid string, r *http.Request, prePage int) ([]Article, pagination.ViewData, error) {
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")

	_pager := pagination.New(r, db, route.Name2URL("categories.show", "id", cid), prePage)

	viewData := _pager.Paging()

	var articles []Article
	err := _pager.Results(&articles)

	return articles, viewData, err
}

func (a *Article) Create() (err error) {
	result := model.DB.Create(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func (a *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&a)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, err
}

func (a *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&a)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, err
}
