package main

import (
	"github.com/attitudefx7/goblog/app/http/middlewares"
	"github.com/attitudefx7/goblog/bootstrap"
	"github.com/attitudefx7/goblog/config"
	c "github.com/attitudefx7/goblog/pkg/config"
	"net/http"
)

func init()  {
	config.Initialize()
}

func main()  {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()


	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
