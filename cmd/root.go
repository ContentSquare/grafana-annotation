package cmd

/*
Copyright © 2020 Contentsquare

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

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	// Singleton logger
	_ "github.com/contentsquare/grafana-annotation/logger"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	verbose bool
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{Use: "grafana-annotation"}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "be more verbose on logs")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config-file", "c", "/etc/grafana-annotation.yaml", "Configuration file")
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogger)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		log.Debugf("loading configuration file %v", cfgFile)
		viper.SetConfigFile(cfgFile)
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "~/.grafana-annotation-poster.yml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("~/.grafana-annotation-poster.yml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger() {
	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Debugf("Setting loglevel to debug.")
	}
}
