/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"fmt"

	"github.com/snivilised/cobrass/src/assistant"
	xi18n "github.com/snivilised/extendio/i18n"
	"golang.org/x/text/language"
)

const (
	AppEmoji        = "🦄"
	ApplicationName = "arcadia"
	RootPsName      = "root-ps"
	SOURCE_ID       = "github.com/snivilised/arcadia"
)

func Execute() {
	bootstrap := Bootstrap{}

	bootstrap.Execute(func(detector LocaleDetector) []string {

		from := xi18n.LoadFrom{
			// Path: "defaults to the exe path",
			// however, you can change this to something
			// else, perhaps you want them to be in ~/your-app/l10n
			// depending on your install process.
			//
			Sources: xi18n.TranslationFiles{
				SOURCE_ID: xi18n.TranslationSource{Name: ApplicationName},
			},
		}

		// read settings from config if they are available there
		// TODO: there is a problem here, config is not
		// read in until after language is setup. This needs to be fixed
		// in another issue.
		//
		err := xi18n.Use(func(uo *xi18n.UseOptions) {
			uo.Tag = detector.Scan()
			uo.From = from
		})

		if err != nil {
			panic(err)
		}

		// TODO: we need to return the real args instead of these
		//
		args := []string{"widget", "-p", "P?<date>", "-t", "30"}
		return args
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
