package bootstrap

import (
	"github.com/attitudefx7/goblog/app/models/article"
	"github.com/attitudefx7/goblog/app/models/category"
	"github.com/attitudefx7/goblog/app/models/user"
	"github.com/attitudefx7/goblog/pkg/config"
	"github.com/attitudefx7/goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB()  {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))

	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))

	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	migration(db)
}

func migration(db *gorm.DB)  {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
		)
}
