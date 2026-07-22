package process

import (
	"fmt"
	"time"
	"gorm.io/gorm"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/database/model"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List tasks with optional filtering by completion status",
	Args:  cobra.MaximumNArgs(1),
	Run:   listTasks,
}

func listTasks(cmd *cobra.Command, args []string) {
	var tasks []model.Task
	var result *gorm.DB
	hideDays := viper.GetInt("behavior.hide_done_after_day")
	if hideDays > 0 {
		cutoff := time.Now().AddDate(0, 0, hideDays)
		result = config.DB.Debug().Where("completed != ? OR completed_at < ?", 1, cutoff.Unix()).Find(&tasks)
	}
	if result.Error != nil {
		fmt.Printf("Proses pengambilan data error karena %d", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println("Tidak ada tugas")
		return
	}

	fmt.Printf("%-4s %-10s %-50s %-10s %s\n", "ID", "STATUS", "DESCRIPTION", "PRIORITY", "CREATED")
	fmt.Println(strings.Repeat("-", 90))

	for _,task := range tasks {
		status := func() string{if task.Completed != 0 {return "SELESAI"}; return "PENDING"}()
		fmt.Printf("%-4d %-10s %-50s %-10s %s\n",
			task.ID,
			status,
			task.Description,
			task.Priority,
			time.Unix(task.CreatedAt, 0))
	}
}
