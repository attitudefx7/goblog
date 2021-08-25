package controllers

import (
	"fmt"
	"github.com/attitudefx7/goblog/app/models/user"
	"github.com/attitudefx7/goblog/app/requests"
	"github.com/attitudefx7/goblog/pkg/auth"
	"github.com/attitudefx7/goblog/pkg/flash"
	"github.com/attitudefx7/goblog/pkg/view"
	"net/http"
)

type AuthController struct {

}

func (auc *AuthController) Register(w http.ResponseWriter, r *http.Request)  {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (auc *AuthController) DoRegister(w http.ResponseWriter, r *http.Request)  {
	_user := user.User{
		Name: r.PostFormValue("name"),
		Email: r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	errs := requests.ValidateRegistrationForm(_user)
	fmt.Print(errs)
	if len(errs) > 0 {
		// 3. 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		// 4. 验证成功，创建数据
		_user.Create()

		if _user.ID > 0 {
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}
}

func (auc *AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	view.RenderSimple(w, view.D{}, "auth.login")
}

func (auc *AuthController) DoLogin(w http.ResponseWriter, r *http.Request)  {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Error": err.Error(),
			"Email": email,
			"Password": password,
		}, "auth.login")
	}
}

func (auc *AuthController) Logout(w http.ResponseWriter, r *http.Request)  {
	auth.Logout()
	flash.Success("您已退出登录")
	http.Redirect(w, r, "/", http.StatusFound)
}
