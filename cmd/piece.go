/*
Copyright © 2025 Andrei Markin<andrey.markin.ppc@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/musgit-dev/musgit/models"
	"github.com/spf13/cobra"
)

var pieceCmd = &cobra.Command{
	Use:   "piece",
	Short: "Working with pieces",
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Composer\tPiece\tPractices")
		for _, piece := range mg.Piece.GetAll() {
			fmt.Fprintf(w, "%s\t%s\t%d\n",
				piece.Composer.Name,
				piece.Name,
				len(piece.Practices))
		}
		w.Flush()
	},
}
var showPieceCmd = &cobra.Command{
	Use:   "show <piece_id>",
	Short: "Shows information about the piece",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pieceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		piece, err := mg.Piece.Get(pieceId)
		if err != nil {
			fmt.Println("Unknown Piece Id:", pieceId)
		}
		displayPiece(piece)
	},
}
var addPieceCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new piece",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		composer := cmd.Flag("composer").Value.String()
		complexity, _ := strconv.Atoi(cmd.Flag("name").Value.String())
		piece, err := mg.Piece.Add(
			name,
			composer,
			models.PieceComplexity(complexity),
		)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Added new piece: ", piece.ID)
		}
	},
}
var practicePieceCmd = &cobra.Command{
	Use:   "practice <piece_id>",
	Short: "Starts practice for a selected piece",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pieceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		lessonId, _ := strconv.ParseInt(
			cmd.Flag("lesson").Value.String(),
			10,
			64,
		)
		practice, err := mg.Practice.Start(pieceId, lessonId)
		if err != nil {
			fmt.Println("Failed to start practice :", err)
		}
		fmt.Println("Started practice: ", practice.ID)
	},
}

func displayPiece(piece models.Piece) {
	fmt.Printf(
		"[%s]: %s\t\tPracticed %d time(s)\n",
		piece.Composer.Name,
		piece.Name,
		len(piece.Practices),
	)

}

func init() {
	rootCmd.AddCommand(pieceCmd)
	pieceCmd.AddCommand(showPieceCmd)
	pieceCmd.AddCommand(addPieceCmd)
	pieceCmd.AddCommand(practicePieceCmd)

	addPieceCmd.Flags().String("name", "", "Name of a piece")
	addPieceCmd.Flags().String("composer", "", "Composer")
	addPieceCmd.Flags().Int64("complexity", 0, "Complexity of a piece")

	practicePieceCmd.Flags().Int64("lesson", 1, "Lesson ID")
}
