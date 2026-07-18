package process

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/model"
)

var CompleteCmd = &cobra.Command{
	Use:   "complete [task-id]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run:   completeTask,
}

func completeTask(cmd *cobra.Command, args []string){
	id, err := strconv.ParseUint(args[0], 0, 8)
	
	var tasks []model.Task
	result := config.DB.Where("id = ?", id).First(&tasks)
	
	if err != nil {
		fmt.Printf("%s not a number type\n", args[0])
		os.Exit(1)
	}

	if result.Error != nil {
		fmt.Printf("Tidak ada data dengan id #%d", id)
	}

	now := time.Now()
	task := model.Task {
		Completed: true,
		CompletedAt: &now,
	}

	ctx := context.Background()
	rows, _ := gorm.G[model.Task](config.DB).Where("id = ?", id).Updates(ctx, task)

	if rows == 0 {
		fmt.Printf("Tidak berhasil mengubah data")
	}
	
	fmt.Printf("Completed task #%d: %s\n", id, tasks[0].Description)
}