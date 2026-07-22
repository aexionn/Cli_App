package process

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/database/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List tasks with optional filtering by completion status",
	Args:  cobra.MaximumNArgs(1),
	Run:   listTasks,
}

func listTasks(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	hideDays := viper.GetInt("behavior.hide_done_after_day")
	if hideDays <= 0 {
		fmt.Println("Atribut tidak valid")
		os.Exit(1)
	}
	cutoff := time.Now().AddDate(0, 0, hideDays)
	result, dbErr := gorm.G[model.Task](config.DB).Where("completed != ? OR completed_at < ?", 1, cutoff.Unix()).Find(ctx)
	if dbErr != nil {
		fmt.Printf("Proses pengambilan data error karena %d", dbErr)
		os.Exit(1)
	}

	if len(result) == 0 {
		fmt.Println("Tidak ada tugas")
		os.Exit(1)
	}

	fmt.Printf("%-4s %-10s %-50s %-10s %s\n", "ID", "STATUS", "DESCRIPTION", "PRIORITY", "CREATED")
	fmt.Println(strings.Repeat("-", 90))

	for _,task := range result {
		status := func() string{if task.Completed != 0 {return "SELESAI"}; return "PENDING"}()
		fmt.Printf("%-4d %-10s %-50s %-10s %s\n",
			task.ID,
			status,
			task.Description,
			task.Priority,
			time.Unix(task.CreatedAt, 0))
	}
}
