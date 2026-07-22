package process

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	allVar "github.com/aexionn/Cli_App/allvar"
	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/database/model"
)

var AddCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long:  "Add a new task with description and optional priority (high, medium, low)",
	Args:  cobra.MinimumNArgs(1),
	Run:   addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	description := strings.Join(args, "")
	

	validatePriorities := map[string]bool{
		"high":   true,
		"medium": true,
		"low":    true,
	}

	if !validatePriorities[allVar.Priority] {
		fmt.Printf("Invalid priority '%s'. Use: high, medium, or low\n", allVar.Priority)
		os.Exit(1)
	}

	task := model.Task {
		Description: description,
		Priority: allVar.Priority,
		Completed: 0,
		CreatedAt: time.Now().Unix(),
	}

	ctx := context.Background()

	result := gorm.WithResult()
	err := gorm.G[model.Task](config.DB, result).Create(ctx, &task)

	if err != nil {
		fmt.Println("Proses memasukkan data gagal")
	}

	fmt.Printf("Added task #%d: %s [%s]\n", task.ID, task.Description, task.Priority)
}