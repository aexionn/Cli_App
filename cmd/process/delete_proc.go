package process

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aexionn/Cli_App/config"
	"github.com/aexionn/Cli_App/database/model"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var deleteStatus string

var DeleteCmd = &cobra.Command{
	Use:   "delete [task-ids...]",
	Short: "Delete tasks",
	Run:   deleteTask,
}

func init() {
	DeleteCmd.Flags().StringVarP(&deleteStatus, "status", "s", "", "Delete tasks by status (e.g., SELESAI, PENDING)")
}

func deleteTask(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	if deleteStatus != "" && len(args) > 0 {
		fmt.Println("Error: Cannot combine specific IDs with status flag")
		os.Exit(1)
	}

	if deleteStatus == "" && len(args) == 0 {
		fmt.Println("Error: Must provide task IDs or --status flag")
		os.Exit(1)
	}

	var targets []model.Task

	if deleteStatus != "" {
		deleteStatus = strings.ToUpper(deleteStatus)
		var completedVal int
		if deleteStatus == "SELESAI" {
			completedVal = 1
		} else if deleteStatus == "PENDING" {
			completedVal = 0
		} else {
			fmt.Printf("Status tidak valid: %s\n", deleteStatus)
			os.Exit(1)
		}

		result, err := gorm.G[model.Task](config.DB).Where("completed = ?", completedVal).Find(ctx)
		if err != nil {
			fmt.Println("Error fetching tasks:", err)
			os.Exit(1)
		}
		targets = result
	} else {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("'%s' is not a valid number\n", arg)
				os.Exit(1)
			}
			ids = append(ids, id)
		}

		result, err := gorm.G[model.Task](config.DB).Where("id IN ?", ids).Find(ctx)
		if err != nil {
			fmt.Println("Error fetching tasks:", err)
			os.Exit(1)
		}
		targets = result
	}

	if len(targets) == 0 {
		fmt.Println("Tidak ada data yang cocok untuk dihapus")
		return
	}

	var targetIDs []byte
	for _, t := range targets {
		targetIDs = append(targetIDs, t.ID)
	}

	rows, _ := gorm.G[model.Task](config.DB).Where("id IN ?", targetIDs).Delete(ctx)

	if rows == 0 {
		fmt.Println("Tidak berhasil menghapus data")
		return
	}

	fmt.Printf("Successfully deleted %d task(s):\n", rows)
	for _, t := range targets {
		fmt.Printf("- Deleted task #%d: %s\n", t.ID, t.Description)
	}
}

