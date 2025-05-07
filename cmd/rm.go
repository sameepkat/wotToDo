/*
Copyright © 2025 sameepkat <sameepsk2@gmail.com>
*/
package cmd

import (
	"bytes"
	"log"
	"strconv"
	"text/template"

	"github.com/sameepkat/wottodo/db"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task",
	Long:  `Deleting a task using id`,
	Run: func(cmd *cobra.Command, args []string) {
		var query string = `DELETE * FROM Tasks WHERE Status="DONE"`
		var ids []int
		if len(args) > 0 {
			for _, arg := range args {
				num, err := strconv.Atoi(arg)
				if err != nil {
					log.Println("error converting " + arg + " to int")
					return
				}
				ids = append(ids, num)
			}

			tmpl, err := template.New("query").Parse(`DELETE FROM Tasks WHERE id IN ({{- range $index, $id := .}} {{if $index}}, {{end}}{{$id}}{{end}})`)
			if err != nil {
				log.Println("Error parsing template: ", err)
				return
			}

			var queryOutput bytes.Buffer
			err = tmpl.Execute(&queryOutput, ids)
			if err != nil {
				log.Println("Error executing template:", err)
				return
			}

			query = queryOutput.String()

			sqlDB, err := db.Exec()
			if err != nil {
				log.Println("error opening connection to database: %w", err)
			}

			db.Remove(sqlDB, query)

		} else {
			log.Println("invalid id")
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("delete", "d", false, "-d [taskid]")
}
