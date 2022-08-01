/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/snivilised/cobrass/src/assistant"
	"golang.org/x/text/language"
)

const AppEmoji = "🦄"
const ApplicationName = "arcadia"

func Execute() {
	bs := bootstrap{}
	bs.execute()
}

type RootParameterSet struct {
	ConfigFile string
	Language   string
	Toggle     bool
}

func setupRootCommand(container *assistant.CobraContainer) {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	const configFlagName = "config"
	const defConfig = ""
	configUsage := fmt.Sprintf("config file (default is $HOME/.%v.yml", ApplicationName)

	root := container.Root()
	paramSet := assistant.NewParamSet[RootParameterSet](root)

	paramSet.BindString(&assistant.FlagInfo{
		Name:               configFlagName,
		Usage:              configUsage,
		Default:            defConfig,
		AlternativeFlagSet: root.PersistentFlags(),
	}, &paramSet.Native.ConfigFile)

	const langFlagName = "lang"
	const defLang = "en-GB"
	const langUsage = "lang defines the language"

	paramSet.BindValidatedString(&assistant.FlagInfo{
		Name:               langFlagName,
		Usage:              langUsage,
		Default:            defLang,
		AlternativeFlagSet: root.PersistentFlags(),
	}, &paramSet.Native.Language, func(value string) error {
		_, err := language.Parse(value)
		return err
	})

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	const toggleShort = "t"
	const toggleUsage = "toggle Help message for toggle"

	paramSet.BindBool(
		assistant.NewFlagInfo(toggleUsage, toggleShort, false),
		&paramSet.Native.Toggle,
	)

	container.MustRegisterParamSet("root-ps", paramSet)
}
