package config

import (
	// "github.com/aexionn/Cli_App/database/model"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	
	DB, err = gorm.Open(sqlite.Open(viper.GetString("storage.path")), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database") 
	}

	// DB.AutoMigrate(&model.Task{}, &model.User{})
	// DB.Migrator().CreateConstraint(&model.User{}, "Tasks")
	// DB.Migrator().CreateConstraint(&model.User{}, "fk_users_tasks")
}