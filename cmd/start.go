/*
Copyright © 2025 Andrei Markin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"musgit/internal/adapters/db"
	"musgit/internal/application/api"
	"strconv"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <piece_id>",
	Short: "Start practice of a specified piece",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pieceId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Incorrect id: %s", args[0])
		}

		err = startPractice(int64(pieceId))
		return err
	},
}

func init() {
	practiceCmd.AddCommand(startCmd)
}

func startPractice(pieceId int64) error {
	dbAdapter, dbErr := db.NewAdapter("musgit.db")
	if dbErr != nil {
		log.Fatalf("Failed to init db, err: %v", dbErr)
	}

	app := api.NewMusgitService(dbAdapter)
	practice, err := app.StartPractice(pieceId)
	log.Printf("Started practice %d for piece %d", practice.ID, pieceId)
	return err
}
