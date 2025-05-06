/*
Copyright © 2025 sameepkat <sameepsk2@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/sameepkat/wottodo/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long:  `Fetch and display all the todos available`,
	Run: func(cmd *cobra.Command, args []string) {
		sqlDB, err := db.Exec()
		if err != nil {
			fmt.Println("error opening database")
			return
		}

		query := "SELECT * FROM Tasks WHERE Status=\"TODO\""
		all, _ := cmd.Flags().GetBool("all")
		todo, _ := cmd.Flags().GetBool("todo")
		done, _ := cmd.Flags().GetBool("done")
		if all {
			query = `SELECT * FROM Tasks`
		}

		if todo {
			query = `SELECT * FROM Tasks WHERE Status="TODO"`
		}

		if done {
			query = `SELECT * FROM Tasks WHERE Status="DONE"`
		}

		res, err := db.List(sqlDB, query)
		if err != nil {
			// fmt.Println("error fetching data: ", err)
			// return
			log.Printf("errror fetching data: ", err)
		}

		// Create variables to store data
		var id uint32
		var title string
		var status string
		var createdAt time.Time

		for res.Next() {
			err := res.Scan(&id, &title, &status, &createdAt)
			if err != nil {
				fmt.Println("error scanning data from table: ", err)
			}
			_, month, day := createdAt.Date()
			min := createdAt.Minute()
			fmt.Printf("%d %s %s %d-%d-%d\n", id, title, status, month, day, min)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Flags
	listCmd.Flags().BoolP("all", "a", false, "List all tasks (TODO, DONE)")
	listCmd.Flags().BoolP("todo", "t", false, "List all taks that have status of TODO")
	listCmd.Flags().BoolP("done", "d", false, "List all taks that have status of DONE")
}
