// /*
// Copyright © 2026 NAME HERE <EMAIL ADDRESS>

// */
// package main

// import "github.com/aexionn/Cli_App/cmd"

// func main() {
// 	// TestInsert()
// 	cmd.Execute()
// }

package main

import (
	"encoding/json"
	"fmt"

	"github.com/aexionn/Cli_App/cmd"
	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/database/model"
	"gorm.io/gorm"
)

func main(){
	config.Connection()
	cmd.Execute()
}

func createUserFromJson(db *gorm.DB, jsonPayload []byte) ( *model.Task, error){
	var task model.Task

	err := json.Unmarshal(jsonPayload, &task)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	result := db.Create(&task)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to insert record: %w", result.Error)
	}

	return &task, nil
}