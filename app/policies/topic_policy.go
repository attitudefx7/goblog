package policies

import (
	"github.com/attitudefx7/goblog/app/models/article"
	"github.com/attitudefx7/goblog/pkg/auth"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserId
}
