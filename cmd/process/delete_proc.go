package process

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/database/model"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [task-id]", 
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run:   deleteTask,
}

func deleteTask (cmd *cobra.Command, args []string) {
	id, err :=  strconv.Atoi(args[0])
	var tasks []model.Task
	result := config.DB.Where("id = ?", id).First(&tasks)
	
	if err != nil {
		fmt.Printf("%s not a number type\n", args[0])
		os.Exit(1)
	}

	if result.Error != nil {
		fmt.Printf("Tidak ada data dengan id #%d", id)
	}


	ctx := context.Background()
	rows, _ := gorm.G[model.Task](config.DB).Where("id = ?", id).Delete(ctx)

	if rows == 0 {
		fmt.Printf("Tidak berhasil menghapus data")
	}
	
	fmt.Printf("Deleted task #%d: %s\n", id, tasks[0].Description)
}