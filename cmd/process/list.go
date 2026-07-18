package process

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	// "gorm.io/gorm"

	// allVar "github.com/aexionn/Cli_App/allvar"
	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/model"
	// task "github.com/aexionn/Cli_App/model"
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
	result := config.DB.Find(&tasks)

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
		status := func() string{if task.Completed != false {return "SELESAI"}; return "PENDING"}()
		fmt.Printf("%-4d %-10s %-50s %-10s %s\n",
			task.ID,
			status,
			task.Description,
			task.Priority,
			task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
