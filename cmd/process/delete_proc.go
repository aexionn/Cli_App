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

func deleteTask(cmd *cobra.Command, args []string) {
	id, convErr := strconv.Atoi(args[0])
	ctx := context.Background()
	result, dbErr := gorm.G[model.Task](config.DB).Where("id = ?", id).First(ctx)

	if convErr != nil {
		fmt.Printf("%s not a number type\n", args[0])
		os.Exit(1)
	}

	if dbErr != nil {
		fmt.Printf("Tidak ada data dengan id #%d", id)
	}

	rows, _ := gorm.G[model.Task](config.DB).Where("id = ?", id).Delete(ctx)

	if rows == 0 {
		fmt.Printf("Tidak berhasil menghapus data")
	}

	fmt.Printf("Deleted task #%d: %s\n", id, result.Description)
}
