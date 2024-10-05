package cmd

import (
	"github.com/kfahad5067/todo-list-with-go/internal/todo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   `add "my new task item"`,
	Short: "add a task",
	Long:  `add a task`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		item := Store.CreateItem(taskName)
		todo.DisplayItems([]todo.DataItem{item}, false)
	},
}
