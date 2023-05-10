package cmd

import (
	"fmt"
	"strings"

	"github.com/AkifhanIlgaz/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Printf("Somethinh went wrong while adding \"%s\" to your task list: %s \n", task, err.Error())
			return
		}

		fmt.Printf("Added \"%s\" to your task list \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
