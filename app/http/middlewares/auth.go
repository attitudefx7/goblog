package middlewares

import (
	"github.com/attitudefx7/goblog/pkg/auth"
	"github.com/attitudefx7/goblog/pkg/flash"
	"net/http"
)

func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)

			return
		}

		next(w, r)
	}
}