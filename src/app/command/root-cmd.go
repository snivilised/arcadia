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

// CLIENT-TODO: define valid properties on the root parameter set
type RootParameterSet struct {
	Language string
}
