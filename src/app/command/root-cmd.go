/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package command

const (
	AppEmoji        = "🦄"
	ApplicationName = "arcadia"
	RootPsName      = "root-ps"
	SourceID        = "github.com/snivilised/arcadia"
)

func Execute() error {
	return (&Bootstrap{}).Root().Execute()
}

// RootParameterSet defines the configuration options exposed on the
// root command's parameter set (CLIENT-TODO: refine these properties).
type RootParameterSet struct {
	// Language defines the IETF BCP 47 language tag.
	Language string
}
