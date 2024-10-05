package cmd

import (
	"fmt"
	"strconv"

	"github.com/kfahad5067/todo-list-with-go/internal/todo"
	"github.com/spf13/cobra"
)

var description string
var isComplete bool
var isInComplete bool

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "Description for the task")
	updateCmd.Flags().BoolVarP(&isComplete, "complete", "c", false, "Mark the task as complete")
	updateCmd.Flags().BoolVarP(&isInComplete, "incomplete", "i", false, "Mark the task as incomplete")

	updateCmd.MarkFlagsOneRequired("description", "complete", "incomplete")
	updateCmd.MarkFlagsMutuallyExclusive("complete", "incomplete")
}

var updateCmd = &cobra.Command{
	Use:   "update <task_id>",
	Short: "update a task",
	Long:  `update a task`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid task_id: %v\n", args[0])
			return
		}

		item := todo.DataItem{}

		if description != "" {
			item.Description = &description
		}
		if isComplete {
			item.IsComplete = &isComplete
		}else if isInComplete {
			isInComplete = false
			item.IsComplete = &isInComplete
		}

		Store.UpdateItem(taskId, item)
	},
}
