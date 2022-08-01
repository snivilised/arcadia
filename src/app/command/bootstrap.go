package command

import (
	"fmt"
	"os"

	"github.com/snivilised/cobrass/src/assistant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type bootstrap struct {
	container *assistant.CobraContainer
}

func (b *bootstrap) execute() {
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
	BuildWidgetCommand(b.container)

	configure(b.container)

	root := b.container.Root()

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func configure(container *assistant.CobraContainer) {
	// This is the functionality previously defined in initConfig, which was
	// invoked as a result of it be passed into cobra.OnInitialize(). This
	// approach was abandoned due to its reliance on global state and the init()
	// function which is an anti-pattern.
	//

	// initConfig reads in config file and ENV variables if set.
	paramSet := container.MustGetParamSet("root-ps").(*assistant.ParamSet[RootParameterSet])

	if paramSet.Native.ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(paramSet.Native.ConfigFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".arcadia" (without extension).
		// NB: 'arcadia' should be renamed as appropriate
		//
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%v", ApplicationName))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
