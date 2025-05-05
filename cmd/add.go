/*
Copyright © 2025 sameepkat <sameepsk2@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/sameepkat/wottodo/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Long:  `Add a new task with a title and optional status.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No arguments Found.")
			cmd.Help()
			return
		}
		title := args[0]

		done, _ := cmd.Flags().GetBool("done")
		status := "TODO"

		sqlDB, err := db.Exec()
		if err != nil {
			fmt.Printf("Failed to add the data: %v", err)
			return
		}

		if done {
			status = "DONE"
		}

		res, err := db.Add(sqlDB, title, status)
		if err != nil {
			fmt.Printf("Failed to add task: %v\n", err)
		} else {
			lastId, err := res.LastInsertId()
			if err != nil {
				fmt.Printf("Could not get last incsert ID: %v\n", err)
			} else {
				fmt.Println(lastId, title, status)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("done", "d", false, "Set Status to DONE")
}
