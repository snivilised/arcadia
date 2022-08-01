/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"
	"os"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const AppEmoji = "🦄"
const ApplicationName = "arcadia"

var cfgFile string
var lang string

var Container = assistant.NewCobraContainer(
	&cobra.Command{
		Use:   "main",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Version: fmt.Sprintf("'%v'", Version),
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	},
)
var rootCommand = Container.Root()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	bindToConfig := &cfgFile
	const configFlagName = "config"
	const defConfig = ""
	const configUsage = "config file (default is $HOME/.arcadia.yml)"
	rootCommand.PersistentFlags().StringVar(bindToConfig, configFlagName, defConfig, configUsage)

	bindToLang := &lang
	const langFlagName = "lang"
	const defLang = "en-GB"
	const langUsage = "lang defines the language"
	rootCommand.PersistentFlags().StringVar(bindToLang, langFlagName, defLang, langUsage)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	const toggleFlagName = "toggle"
	const toggleShort = "t"
	const toggleUsage = "Help message for toggle"
	rootCommand.Flags().BoolP(toggleFlagName, toggleShort, false, toggleUsage)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".main" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".main")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
