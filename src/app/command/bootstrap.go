package command

import (
	"fmt"
	"os"

	"github.com/cubiest/jibberjabber"
	"github.com/samber/lo"
	"github.com/snivilised/arcadia/src/locale"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/snivilised/cobrass/src/assistant/configuration"
	ci18n "github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/li18ngo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

type LocaleDetector interface {
	Scan() language.Tag
}

// Jabber is a LocaleDetector implemented using jibberjabber.
type Jabber struct {
}

// Scan returns the detected language tag.
func (j *Jabber) Scan() language.Tag {
	lang, _ := jibberjabber.DetectIETF()
	return language.MustParse(lang)
}

type ConfigInfo struct {
	Name       string
	ConfigType string
	ConfigPath string
	Viper      configuration.ViperConfig
}

type ConfigureOptions struct {
	Detector LocaleDetector
	Config   ConfigInfo
}

type ConfigureOptionFn func(*ConfigureOptions)

// Bootstrap represents construct that performs start up of the cli
// without resorting to the use of Go's init() mechanism and minimal
// use of package global variables.
type Bootstrap struct {
	container *assistant.CobraContainer
	options   ConfigureOptions
}

func (b *Bootstrap) prepare() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	b.options = ConfigureOptions{
		Detector: &Jabber{},
		Config: ConfigInfo{
			Name:       ApplicationName,
			ConfigType: "yaml",
			ConfigPath: home,
			Viper:      &configuration.GlobalViperConfig{},
		},
	}
}

// Root builds the command tree and returns the root command, ready
// to be executed.
func (b *Bootstrap) Root(options ...ConfigureOptionFn) *cobra.Command {
	b.prepare()

	for _, fo := range options {
		fo(&b.options)
	}

	b.configure()

	// all these string literals here should be made translate-able
	//

	b.container = assistant.NewCobraContainer(
		&cobra.Command{
			Use:     "main",
			Short:   li18ngo.Text(locale.RootCmdShortDescTemplData{}),
			Long:    li18ngo.Text(locale.RootCmdLongDescTemplData{}),
			Version: fmt.Sprintf("'%v'", Version),
			// Uncomment the following line if your bare application
			// has an action associated with it:
			// Run: func(cmd *cobra.Command, args []string) { },
		},
	)

	b.buildRootCommand(b.container)
	b.buildWidgetCommand(b.container)

	return b.container.Root()
}

func (b *Bootstrap) configure() {
	vc := b.options.Config.Viper
	ci := b.options.Config

	vc.SetConfigName(ci.Name)
	vc.SetConfigType(ci.ConfigType)
	vc.AddConfigPath(ci.ConfigPath)
	vc.AutomaticEnv()

	err := vc.ReadInConfig()

	handleLangSetting()

	if err != nil {
		msg := li18ngo.Text(locale.UsingConfigFileTemplData{
			ConfigFileName: viper.ConfigFileUsed(),
		})
		fmt.Fprintln(os.Stderr, msg)
	}
}

func handleLangSetting() {
	tag := lo.TernaryF(viper.InConfig("lang"),
		func() language.Tag {
			lang := viper.GetString("lang")
			parsedTag, err := language.Parse(lang)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			return parsedTag
		},
		func() language.Tag {
			return li18ngo.DefaultLanguage
		},
	)

	err := li18ngo.Use(func(uo *li18ngo.UseOptions) {
		uo.Tag = tag
		uo.From = li18ngo.LoadFrom{
			Sources: li18ngo.TranslationFiles{
				SourceID: li18ngo.TranslationSource{Name: ApplicationName},

				// By adding in the source for cobrass, we relieve the client from having
				// to do this. After-all, it should be taken as read that since any
				// instantiation of arcadia (ie a project using this template) is by
				// necessity dependent on cobrass, it's source should be loaded so that a
				// localizer can be created for it.
				//
				// The client app has to make sure that when their app is deployed,
				// the translations file(s) for cobrass are named as 'cobrass', as you
				// can see below, that is the name assigned to the app name of the
				// source.
				//
				ci18n.CobrassSourceID: li18ngo.TranslationSource{Name: "cobrass"},
			},
		}
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (b *Bootstrap) buildRootCommand(container *assistant.CobraContainer) {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//
	root := container.Root()
	paramSet := assistant.NewParamSet[RootParameterSet](root)

	paramSet.BindValidatedString(&assistant.FlagInfo{
		Name:               "lang",
		Usage:              li18ngo.Text(locale.RootCmdLangUsageTemplData{}),
		Default:            li18ngo.DefaultLanguage.String(),
		AlternativeFlagSet: root.PersistentFlags(),
	}, &paramSet.Native.Language, func(value string, _ *pflag.Flag) error {
		_, err := language.Parse(value)
		return err
	})

	container.MustRegisterParamSet(RootPsName, paramSet)
}
