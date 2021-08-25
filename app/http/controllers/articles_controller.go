package controllers

import (
	"fmt"
	article2 "github.com/attitudefx7/goblog/app/models/article"
	"github.com/attitudefx7/goblog/app/policies"
	"github.com/attitudefx7/goblog/app/requests"
	"github.com/attitudefx7/goblog/pkg/auth"
	"github.com/attitudefx7/goblog/pkg/logger"
	"github.com/attitudefx7/goblog/pkg/route"
	"github.com/attitudefx7/goblog/pkg/view"
	"net/http"
)

type ArticlesController struct {
	BaseController
}

func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request)  {
	articles, pagerData, err := article2.GetAll(r, 2)
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request)  {
	// 获取参数
	id := route.GetRouteVariable("id", r)

	// 查询文章
	article, err := article2.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Article": article,
		},"articles.show", "articles._article_meta")
	}
}

func (ac *ArticlesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (ac *ArticlesController) Store(w http.ResponseWriter, r *http.Request)  {
	_article := article2.Article{
		Title: r.PostFormValue("title"),
		Body: r.PostFormValue("body"),
		UserId: auth.User().ID,
	}

	err := requests.ValidateArticleForm(_article)

	if len(err) == 0 {

		_article.Create()
		if _article.ID > 0 {
			showURL := route.Name2URL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, showURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}

	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors": err,
		}, "articles.create", "articles._form_field")
	}
}


func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request)  {

	// 获取参数
	id := route.GetRouteVariable("id", r)

	// 查询是否存在
	article, err := article2.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			view.Render(w, view.D{
				"Article": article,
				"Errors": view.D{},
			}, "articles.edit", "articles._form_field")
		}
	}
}

func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouteVariable("id", r)

	_article, err := article2.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")

			errors := requests.ValidateArticleForm(_article)

			if len(errors) == 0 {

				rowsAffected, err := _article.Update()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 服务器内部错误")
				}

				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何更改")
				}

			} else {
				view.Render(w, view.D{
					"Article": _article,
					"Errors":  errors,
				}, "articles.edit", "articles._form_field")
			}
		}
	}
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouteVariable("id", r)

	article, err := article2.Get(id)
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			rowsAffected, err := article.Delete()

			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			} else {
				if rowsAffected > 0 {
					indexURL := route.Name2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}
	}
}
