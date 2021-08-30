package controllers

import (
	"fmt"
	"net/http"
)

type PagesController struct {

}

func (pc *PagesController) Home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 GoBlog!</h1>")
}


func (pc *PagesController) About(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:attitudefx7@gmail.com\">attitudefx7@gmail.com</a>")
}

func (pc *PagesController) NotFound(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}
