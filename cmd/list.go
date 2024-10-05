package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kfahad5067/todo-list-with-go/internal/todo"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks",
	Long:  `list the tasks`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		items := Store.ReadItems()
		todo.DisplayItems(items)	
	},
}
