package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete <task_id>",
	Short: "delete a task",
	Long:  `delete a task`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid task_id: %v\n", args[0])
			return
		}

		Store.DeleteItem(taskId)
	},
}
