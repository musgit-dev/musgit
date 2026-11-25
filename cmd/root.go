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
	"log/slog"
	"os"
	"strings"

	"github.com/musgit-dev/musgit/internal/musgit"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbUri string
var configFile string
var mg *musgit.Musgit

var rootCmd = &cobra.Command{
	Use:     "mg",
	Short:   "Musgit CLI client.",
	Version: "0.0.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		dbUri := viper.GetString("db-uri")
		if dbUri == "" {
			return fmt.Errorf("Missing --db-uri flag.")
		}
		mg = musgit.New(dbUri)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().
		StringVarP(&dbUri, "db-uri", "d", "", "DB with data")
	rootCmd.PersistentFlags().
		StringVarP(&configFile, "config", "c", "",
			"Config file (Default is $HOME/.musgit/musgit.yaml")

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("MUSGIT")

	err := viper.BindPFlag("db-uri", rootCmd.PersistentFlags().Lookup("db-uri"))
	if err != nil {
		fmt.Println("DB URI is not provided.")
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home + "/.musgit")
		viper.SetConfigName("musgit")
		viper.AutomaticEnv()

		_ = viper.ReadInConfig()
		slog.Info("config:", "path", viper.ConfigFileUsed())
	}

}
