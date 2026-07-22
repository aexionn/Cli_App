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
	"github.com/aexionn/Cli_App/database/model"
)

var CompleteCmd = &cobra.Command{
	Use:   "complete [task-id]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run:   completeTask,
}

func completeTask(cmd *cobra.Command, args []string){
	id, convErr := strconv.ParseUint(args[0], 0, 8)
	ctx := context.Background()
	result, dbErr := gorm.G[model.Task](config.DB).Where("id = ?", id).First(ctx)
	
	if convErr != nil {
		fmt.Printf("%s not a number type\n", args[0])
		os.Exit(1)
	}

	if dbErr != nil {
		fmt.Printf("Tidak ada data dengan id #%d", id)
	}

	now := time.Now()
	task := model.Task {
		Completed: 1,
		CompletedAt: now.Unix(),
	}

	rows, _ := gorm.G[model.Task](config.DB).Where("id = ?", id).Updates(ctx, task)

	if rows == 0 {
		fmt.Printf("Tidak berhasil mengubah data")
	}
	
	fmt.Printf("Completed task #%d: %s\n", id, result.Description)
}