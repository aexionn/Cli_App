package cmd

import (
	"os"

	"github.com/spf13/cobra"

	allVar "github.com/aexionn/Cli_App/allvar"
	config "github.com/aexionn/Cli_App/config"
	proc "github.com/aexionn/Cli_App/cmd/process"
)

var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "A personal task manager",
	Long: `taskman is a CLI task manager that helps you organize your work.
Store tasks locally with priorities, mark them complete, and keep
track of your productivity over time.`,
	// PersistentPreRun:  loadTasks,
	// PersistentPostRun: saveTasks,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&allVar.TaskFile, "file", "", "task file (default default is $HOME/Documents/.taskman.json)")
	proc.AddCmd.Flags().StringVarP(&allVar.Priority, "priority", "p", "medium", "task priority (high, medium, low)")
	proc.ListCmd.Flags().BoolVarP(&allVar.ShowAll, "all", "a", false, "show completed tasks too")

	rootCmd.AddCommand(proc.AddCmd, proc.ListCmd, proc.CompleteCmd, proc.DeleteCmd)
	config.SetupConfig()
}
