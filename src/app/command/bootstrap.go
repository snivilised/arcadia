package command

import (
	"bytes"
	"fmt"
	"os"

	"github.com/cubiest/jibberjabber"
	"github.com/snivilised/arcadia/src/internal/translate"
	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

type LocaleDetector interface {
	Scan() language.Tag
}

// Jabber is a LocaleDetector implemented using jibberjabber
//
type Jabber struct {
}

// Scan returns the detected language tag
//
func (j *Jabber) Scan() language.Tag {
	lang, _ := jibberjabber.DetectIETF()
	return language.MustParse(lang)
}

// Bootstrap represents construct that performs start up of the cli without resorting to
// the use of Go's init() mechanism and minimal use of package global variables
//
type Bootstrap struct {
	Detector  LocaleDetector
	container *assistant.CobraContainer
}

// Execute runs the bootstrap. This is typically invoked from the root command, which
// typically initialises the translate package.
//
func (b *Bootstrap) Execute(initialise func(LocaleDetector) []string) {

	if b.Detector == nil {
		b.Detector = &Jabber{}
	}
	args := initialise(b.Detector)
	configure()

	// all these string literals here should be made translateable
	//
	b.container = assistant.NewCobraContainer(
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

	setupRootCommand(b.container)
	buildWidgetCommand(b.container)

	root := b.container.Root()
	var err error

	if args != nil {
		buf := new(bytes.Buffer)
		root.SetOut(buf)
		root.SetErr(buf)
		root.SetArgs(args)

		_, err = root.ExecuteC()
	} else {
		err = root.Execute()
	}

	if err != nil {
		os.Exit(1)
	}
}

type configureOptions struct {
	confileFile *string
}

type ConfigureOptionFn func(*configureOptions)

func configure(options ...ConfigureOptionFn) {
	var configFile string

	o := configureOptions{
		confileFile: &configFile,
	}
	for _, fo := range options {
		fo(&o)
	}

	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".arcadia" (without extension).
		// NB: 'arcadia' should be renamed as appropriate
		//
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%v", APPLICATION_NAME))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	handleLangSetting()
}

func handleLangSetting() {
	if viper.InConfig("lang") {
		lang := viper.GetString("lang")
		tag, err := language.Parse(lang)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_ = translate.UseTag(tag)
	}
}
