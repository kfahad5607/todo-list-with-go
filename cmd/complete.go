package cmd

import (
	"fmt"
	"strconv"

	"github.com/kfahad5067/todo-list-with-go/internal/todo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete <task_id>",
	Short: "Mark a task as complete",
	Long:  `Mark a task as complete`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid task_id: %v\n", args[0])
			return
		}
		
		isComplete := true
		Store.UpdateItem(taskId, todo.DataItem{IsComplete: &isComplete})
	},
}
