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

var lessonCmd = &cobra.Command{
	Use:   "lesson",
	Short: "Works with lessons",
	RunE: func(cmd *cobra.Command, args []string) error {
		lessons := mg.Lesson.GetAll()
		for i, lesson := range lessons {
			fmt.Printf("%d: \t%d\n", i, lesson.ID)
		}
		return nil

	},
}

var startLessonCmd = &cobra.Command{
	Use:   "start",
	Short: "Stars new lesson",
	Run: func(cmd *cobra.Command, args []string) {
		lesson, _ := mg.Lesson.Start()
		fmt.Println("lesson started:", lesson.ID)
	},
}

var pauseLessonCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses current lesson",
	Run: func(cmd *cobra.Command, args []string) {
		err := mg.Lesson.PauseCurrent()
		if err != nil {
			fmt.Println("Failed to pause lesson:", err)
		}

		fmt.Println("lesson paused")
	},
}

var resumeLessonCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes current lesson",
	Run: func(cmd *cobra.Command, args []string) {
		err := mg.Lesson.ResumeCurrent()
		if err != nil {
			fmt.Println("Failed to resume lesson:", err)
		}

		fmt.Println("lesson resumed")
	},
}
var stopLessonCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops current lesson",
	Run: func(cmd *cobra.Command, args []string) {
		err := mg.Lesson.StopCurrent()
		if err != nil {
			fmt.Println("Failed to stop lesson:", err)
		}

		fmt.Println("lesson stopped")
	},
}

func init() {
	rootCmd.AddCommand(lessonCmd)
	lessonCmd.AddCommand(startLessonCmd)
	lessonCmd.AddCommand(resumeLessonCmd)
	lessonCmd.AddCommand(stopLessonCmd)
	lessonCmd.AddCommand(pauseLessonCmd)
}
