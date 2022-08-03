/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/snivilised/arcadia/src/internal/translate"
	"github.com/snivilised/cobrass/src/assistant"
	"golang.org/x/text/language"
)

const APP_EMOJI = "🦄"
const APPLICATION_NAME = "arcadia"
const ROOT_PSNAME = "root-ps"

func Execute() {
	bs := Bootstrap{}
	bs.Execute(func(detector LocaleDetector) []string {
		translate.Initialise(func(o *translate.LanguageInitOptions) {
			o.Detected = detector.Scan()
			o.App = APPLICATION_NAME
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

	root := container.Root()
	paramSet := assistant.NewParamSet[RootParameterSet](root)

	paramSet.BindString(&assistant.FlagInfo{
		Name:               "config",
		Usage:              fmt.Sprintf("config file (default is $HOME/.%v.yml", APPLICATION_NAME),
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
	const TOGGLE_SHORT = "t"
	const TOGGLE_USAGE = "toggle Help message for toggle"

	paramSet.BindBool(
		assistant.NewFlagInfo(TOGGLE_USAGE, TOGGLE_SHORT, false),
		&paramSet.Native.Toggle,
	)

	container.MustRegisterParamSet(ROOT_PSNAME, paramSet)
}
