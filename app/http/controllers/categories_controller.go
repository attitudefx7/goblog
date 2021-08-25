package controllers

import (
	"fmt"
	"github.com/attitudefx7/goblog/app/models/article"
	"github.com/attitudefx7/goblog/app/models/category"
	"github.com/attitudefx7/goblog/app/requests"
	"github.com/attitudefx7/goblog/pkg/logger"
	"github.com/attitudefx7/goblog/pkg/route"
	"github.com/attitudefx7/goblog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

func (cc CategoriesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.Render(w, view.D{}, "categories.create")
}

func (cc CategoriesController) Store(w http.ResponseWriter, r *http.Request)  {
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	errors := requests.ValidateCategoryForm(_category)

	if len(errors) == 0 {
		_category.Create()
		if _category.ID > 0 {
			fmt.Fprint(w, "创建成功！")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors": errors,
		}, "categories.create")
	}
}

func (cc CategoriesController) Show(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouteVariable("id", r)

	_category, err := category.Get(id)
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 分类未找到")
	}

	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 2)

	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}
