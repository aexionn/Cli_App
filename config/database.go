package config

import (
	// "github.com/aexionn/Cli_App/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("my_database"), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database") 
	}

	// DB.AutoMigrate(&model.Task{}, &model.User{})
	// DB.Migrator().CreateConstraint(&model.User{}, "Tasks")
	// DB.Migrator().CreateConstraint(&model.User{}, "fk_users_tasks")
}