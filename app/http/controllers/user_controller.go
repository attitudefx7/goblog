package controllers

import (
	"fmt"
	"github.com/attitudefx7/goblog/app/models/article"
	"github.com/attitudefx7/goblog/app/models/user"
	"github.com/attitudefx7/goblog/pkg/logger"
	"github.com/attitudefx7/goblog/pkg/route"
	"github.com/attitudefx7/goblog/pkg/view"
	"net/http"
)

type UserController struct {
	BaseController
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request)  {
	id := route.GetRouteVariable("id", r)

	_user, err := user.Get(id)

	if err != nil {
		uc.ResponseForSQLError(w, err)
	} else {
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
