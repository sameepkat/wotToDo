/*
Copyright © 2025 sameepkat <sameepsk2@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sameepkat/wottodo/db"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long: `Updating a task using id
Examples:
		wottodo update 1 -t
		wottodo update 2 -s
		wottodo update 5 -ts	`,
	Run: func(cmd *cobra.Command, args []string) {
		t, _ := cmd.Flags().GetBool("title")
		s, _ := cmd.Flags().GetBool("status")

		if len(args) == 0 {
			log.Println("Please provide a task ID.")
			return
		}
		id := args[0]

		sqlDb, err := db.Exec()
		if err != nil {
			log.Println("error finding data with that id: ", err)
			return
		}

		row := sqlDb.QueryRow("SELECT id, title, status FROM Tasks WHERE id = ?", id)

		var taskId int
		var oldTitle, oldStatus string
		err = row.Scan(&taskId, &oldTitle, &oldStatus)

		if err != nil {
			log.Println("task with that id not found")
			return
		}

		var newTitle = oldTitle
		var newStatus = oldStatus

		if t {
			newTitle = askUser("Enter new title: ")
		}
		if s {
			newStatus = askUser("Enter new status: ")
		}

		_, err = sqlDb.Exec("UPDATE tasks SET title = ?, status = ? WHERE id = ?", newTitle, newStatus, taskId)

		if err != nil {
			log.Println("Error updating task:", err)
			return
		}

		fmt.Println("Task updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolP("title", "t", false, "Update title")
	updateCmd.Flags().BoolP("status", "s", false, "Update status")
}

func askUser(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(input)
	return userInput

}
