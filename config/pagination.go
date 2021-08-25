package config

import "github.com/attitudefx7/goblog/pkg/config"

func init()  {
	config.Add("pagination", config.StrMap{
		"prepage": 10,

		"url_query": "page",
	})
}
