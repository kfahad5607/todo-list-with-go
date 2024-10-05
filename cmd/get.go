package cmd

import (
	"fmt"
	"strconv"

	"github.com/kfahad5067/todo-list-with-go/internal/todo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get <task_id>",
	Short: "get a task",
	Long:  `get a task`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid task_id: %v\n", args[0])
			return
		}

		item := Store.ReadItem(taskId)
		todo.DisplayItems([]todo.DataItem{item}, false)
	},
}
