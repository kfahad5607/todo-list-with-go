package cmd

import (
	"github.com/kfahad5067/todo-list-with-go/internal/todo"
	"github.com/spf13/cobra"
)


// var Store todo.DataStore = todo.NewCsvStore("data")
// var Store todo.DataStore = todo.NewJsonStore("data")
var Store todo.DataStore = todo.NewSqliteStore("data")

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
