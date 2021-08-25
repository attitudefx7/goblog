package requests

import (
	"github.com/attitudefx7/goblog/app/models/category"
	"github.com/thedevsaddam/govalidator"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称必填",
			"min_cn:名称长度需大于 2 个字",
			"max_cn:名称长度最大为 8 个字",
		},
	}

	opts := govalidator.Options{
		Data: &data,
		Rules: rules,
		Messages: messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
