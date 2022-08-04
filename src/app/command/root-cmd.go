/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/snivilised/cobrass/src/assistant"
	"golang.org/x/text/language"

	"github.com/snivilised/arcadia/src/internal/translate"
)

const AppEmoji = "🦄"
const ApplicationName = "arcadia"
const RootPsName = "root-ps"

func Execute() {
	bs := Bootstrap{}
	bs.Execute(func(detector LocaleDetector) []string {
		translate.Initialise(func(o *translate.LanguageInitOptions) {
			o.Detected = detector.Scan()
			o.App = ApplicationName
		})
		return nil
	})
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
	//
	root := container.Root()
	paramSet := assistant.NewParamSet[RootParameterSet](root)

	paramSet.BindString(&assistant.FlagInfo{
		Name:               "config",
		Usage:              fmt.Sprintf("config file (default is $HOME/.%v.yml", ApplicationName),
		Default:            "",
		AlternativeFlagSet: root.PersistentFlags(),
	}, &paramSet.Native.ConfigFile)

	paramSet.BindValidatedString(&assistant.FlagInfo{
		Name:               "lang",
		Usage:              "lang defines the language",
		Default:            "en-GB",
		AlternativeFlagSet: root.PersistentFlags(),
	}, &paramSet.Native.Language, func(value string) error {
		_, err := language.Parse(value)
		return err
	})

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	const (
		ToggleShort = "t"
		ToggleUsage = "toggle Help message for toggle"
	)

	paramSet.BindBool(
		assistant.NewFlagInfo(ToggleUsage, ToggleShort, false),
		&paramSet.Native.Toggle,
	)

	container.MustRegisterParamSet(RootPsName, paramSet)
}
