package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	allVar "github.com/aexionn/Cli_App/allvar"
	proc "github.com/aexionn/Cli_App/cmd/process"
)

var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "A personal task manager",
	Long: `taskman is a CLI task manager that helps you organize your work.

Store tasks locally with priorities, mark them complete, and keep
track of your productivity over time.`,
	PersistentPreRun:  loadTasks,
	PersistentPostRun: saveTasks,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&allVar.TaskFile, "file", "", "task file (default default is $HOME/Documents/.taskman.json)")
	proc.AddCmd.Flags().StringVarP(&allVar.Priority, "priority", "p", "medium", "task priority (high, medium, low)")
	proc.ListCmd.Flags().BoolVarP(&allVar.ShowAll, "all", "a", false, "show completed tasks too")
	
	rootCmd.AddCommand(proc.AddCmd, proc.ListCmd, proc.CompleteCmd, proc.DeleteCmd)
	setupConfig()
}

func setupConfig() {
	viper.SetConfigName("taskman")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.SetDefault("priority", "medium")
	viper.SetDefault("file", filepath.Join(os.Getenv("HOME/Documents"), ".taskman.json"))

	viper.ReadInConfig()
}

func getTaskFile() string {
	if allVar.TaskFile != "" {
		return allVar.TaskFile
	}
	return viper.GetString("file")
}

func loadTasks(cmd *cobra.Command, args []string) {
	file := getTaskFile()
	allVar.TaskManagerVar = &allVar.TaskManager{
		Tasks: make([]allVar.Task, 0),
		NextID: 1,
		FilePath: file,
	}

	if data, err := os.ReadFile(file); err == nil {
		json.Unmarshal(data, allVar.TaskManagerVar)
	}
}

func saveTasks(cmd *cobra.Command, args []string) {
	data, err := json.MarshalIndent(allVar.TaskManagerVar, "", " ")
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	os.WriteFile(allVar.TaskManagerVar.FilePath, data, 0644)
}