package domain

import "github.com/snivilised/cobrass/src/assistant"

// OutputFormatEnum is a placeholder enumeration that represents the
// different output formats supported by the widget command.
// CLIENT-TODO: remove or replace this dummy enum definition as needed.
type OutputFormatEnum int

const (
	_ OutputFormatEnum = iota
	XMLFormatEn
	JSONFormatEn
	TextFormatEn
	ScribbleFormatEn
)

var OutputFormatEnumInfo = assistant.NewEnumInfo(assistant.AcceptableEnumValues[OutputFormatEnum]{
	XMLFormatEn:      []string{"xml", "x"},
	JSONFormatEn:     []string{"json", "j"},
	TextFormatEn:     []string{"text", "tx"},
	ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
})

// WidgetParameterSet defines the parameters that control the behaviour
// of the widget command, including input location, formatting and
// filtering options.
type WidgetParameterSet struct {
	// Directory to process.
	Directory string
	// Concise output flag.
	Concise bool
	// Pattern to match.
	Pattern string
	// Threshold for processing.
	Threshold uint

	// Format of the output.
	Format OutputFormatEnum
	// FormatEn is the enumerable output format.
	FormatEn assistant.EnumValue[OutputFormatEnum]
}

// WidgetParamSetPtr is a convenience alias for a parameter set that
// wraps a WidgetParameterSet instance.
type WidgetParamSetPtr = *assistant.ParamSet[WidgetParameterSet]
