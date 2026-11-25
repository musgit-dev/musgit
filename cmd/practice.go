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
	"strconv"

	"github.com/spf13/cobra"
)

var practiceCmd = &cobra.Command{
	Use:   "practice",
	Short: "Shows information on practices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("practice called")
	},
}
var startPracticeCmd = &cobra.Command{
	Use:   "start <piece_id>",
	Short: "Starts practice for a piece",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pieceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		_, err = mg.Practice.Start(pieceId, 0)
		if err != nil {
			fmt.Println(err)
		}
	},
}
var stopPracticeCmd = &cobra.Command{
	Use:   "stop <piece_id>",
	Short: "Stops practice for a piece",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pieceId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		_, err = mg.Practice.Stop(pieceId)
		if err != nil {
			fmt.Println(err)
		}
	},
}
var evalutePracticeCmd = &cobra.Command{
	Use:   "evalute <practice_id> <evaluation>",
	Short: "Provide evaluation for the practice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("practice evaluated")
	},
}

func init() {
	rootCmd.AddCommand(practiceCmd)
	practiceCmd.AddCommand(startPracticeCmd)
	practiceCmd.AddCommand(stopPracticeCmd)
	practiceCmd.AddCommand(evalutePracticeCmd)
}
