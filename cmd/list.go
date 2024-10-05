package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kfahad5067/todo-list-with-go/internal/todo"
)

var showAll bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Display all tasks, including completed")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks",
	Long:  `list the tasks`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		items := Store.ReadItems(showAll)
		todo.DisplayItems(items)
	},
}
