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

	"github.com/spf13/cobra"
)

var warmupCmd = &cobra.Command{
	Use:   "warmup",
	Short: "Starts warmup",
	Run: func(cmd *cobra.Command, args []string) {
		warmup, err := mg.Practice.Warmup(0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Started warmup", warmup.ID)
	},
}
var stopWarmupCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops warmup",
	Run: func(cmd *cobra.Command, args []string) {
		warmup, err := mg.Practice.StopWarmup()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Stopped warmup", warmup.ID)
	},
}

func init() {
	rootCmd.AddCommand(warmupCmd)
	warmupCmd.AddCommand(stopWarmupCmd)
}
