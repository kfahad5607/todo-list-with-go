package cmd

import (
	"github.com/kfahad5067/todo-list-with-go/internal/todo"
	"github.com/spf13/cobra"
)


// var Store todo.DataStore = todo.CsvStore{ StoreName: "data.csv"}
var Store todo.DataStore = todo.JsonStore{ StoreName: "data.json"}

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A todo list for the terminal",
	Long:  `A todo list for the terminal`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
