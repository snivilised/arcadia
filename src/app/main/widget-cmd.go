package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/snivilised/cobrass/src/adapters"
	"github.com/spf13/cobra"
)

type OutputFormatEnum int

const (
	_ OutputFormatEnum = iota
	XmlFormatEn
	JsonFormatEn
	TextFormatEn
	ScribbleFormatEn
)

var OutputFormatEnumInfo = adapters.NewEnumInfo(adapters.AcceptableEnumValues[OutputFormatEnum]{
	XmlFormatEn:      []string{"xml", "x"},
	JsonFormatEn:     []string{"json", "j"},
	TextFormatEn:     []string{"text", "tx"},
	ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
})
var OutputFormatEn adapters.EnumValue[OutputFormatEnum]

type WidgetParameterSet struct {
	Directory string
	Format    OutputFormatEnum
	Concise   bool
	Pattern   string
	Threshold uint
}

func init() {
	widgetCommand := &cobra.Command{
		Use:   "widget",
		Short: "widget sub command",
		Long:  "Long description of the widget command",
		RunE: func(cmd *cobra.Command, args []string) error {
			var appErr error = nil

			ps := Container.MustGetParamSet("widget-ps").(*adapters.ParamSet[WidgetParameterSet])

			if err := ps.Validate(); err == nil {
				native := ps.Native

				// rebind enum into native member
				//
				native.Format = OutputFormatEn.Value()

				// optionally invoke cross field validation
				//
				if xv := ps.CrossValidate(func(ps *WidgetParameterSet) error {
					condition := (ps.Format == XmlFormatEn)
					if condition {
						return nil
					}
					return fmt.Errorf("format: '%v' is invalid", ps.Format)
				}); xv == nil {
					fmt.Printf("%v %v Running widget\n", AppEmoji, ApplicationName)
					// ---> execute application core with the parameter set (native)
					//
					// appErr = runApplication(native)
					//
				} else {
					return xv
				}
			} else {
				return err
			}

			return appErr
		},
	}

	paramSet := adapters.NewParamSet[WidgetParameterSet](widgetCommand)
	paramSet.BindValidatedString(
		adapters.NewFlagInfo("directory", "d", "/foo-bar"),
		&paramSet.Native.Directory,
		func(value string) error {
			if _, err := os.Stat(value); err != nil {
				if os.IsNotExist(err) {
					return err
				}
			}
			return nil
		},
	)

	OutputFormatEn = OutputFormatEnumInfo.NewValue()
	paramSet.BindValidatedEnum(
		adapters.NewFlagInfo("format", "f", "xml"),
		&OutputFormatEn.Source,
		func(value string) error {
			if OutputFormatEnumInfo.En(value) == XmlFormatEn {
				return nil
			}
			return fmt.Errorf("only xml format is currently supported, other formats available in future release")
		},
	)

	paramSet.BindBool(
		adapters.NewFlagInfo("concise", "c", false),
		&paramSet.Native.Concise,
	)

	paramSet.BindValidatedString(
		adapters.NewFlagInfo("pattern", "p", ""),
		&paramSet.Native.Pattern,
		func(value string) error {
			result := strings.Contains(value, "P?<date>") ||
				(strings.Contains(value, "P?<d>") && strings.Contains(value, "P?<m>") &&
					strings.Contains(value, "P?<m>"))

			if result {
				return nil
			}

			return fmt.Errorf("pattern is invalid, missing mandatory capture groups ('date' or 'd', 'm', and 'y')")
		},
	)
	widgetCommand.MarkFlagRequired("pattern")

	const lo = uint(25)
	const hi = uint(50)

	paramSet.BindValidatedUintWithin(
		adapters.NewFlagInfo("threshold", "t", uint(10)),
		&paramSet.Native.Threshold,
		lo, hi,
	)

	Container.MustRegisterRootedCommand(widgetCommand)
	Container.MustRegisterParamSet("widget-ps", paramSet)
}
