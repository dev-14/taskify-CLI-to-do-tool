package cmd

import (
	"fmt"
	"strings"
	"taskify/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your tasklist",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTasks(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
