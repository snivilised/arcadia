package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/snivilised/cobrass/src/assistant"
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

type WidgetParameterSet struct {
	Directory string
	Format    OutputFormatEnum
	Concise   bool
	Pattern   string
	Threshold uint

	// the following are supporting fields required for widget command
	//
	OutputFormatEnumInfo *assistant.EnumInfo[OutputFormatEnum]
	OutputFormatEn       assistant.EnumValue[OutputFormatEnum]
}

const WIDGET_PSNAME = "widget-ps"

func buildWidgetCommand(container *assistant.CobraContainer) *cobra.Command {
	// to test: arcadia widget -d ./some-existing-file -p "P?<date>" -t 30
	//
	widgetCommand := &cobra.Command{
		Use:   "widget",
		Short: "widget sub command",
		Long:  "Long description of the widget command",
		RunE: func(cmd *cobra.Command, args []string) error {
			var appErr error = nil

			// check for alternative config file setting
			//
			if rps := container.MustGetParamSet(ROOT_PSNAME).(*assistant.ParamSet[RootParameterSet]); rps.Native.ConfigFile != "" {
				configure(func(co *configureOptions) {
					*co.confileFile = rps.Native.ConfigFile
				})
			}

			ps := container.MustGetParamSet(WIDGET_PSNAME).(*assistant.ParamSet[WidgetParameterSet])

			if err := ps.Validate(); err == nil {
				native := ps.Native

				// rebind enum into native member
				//
				native.Format = native.OutputFormatEn.Value()

				// optionally invoke cross field validation
				//
				if xv := ps.CrossValidate(func(ps *WidgetParameterSet) error {
					condition := (ps.Format == XmlFormatEn)
					if condition {
						return nil
					}
					return fmt.Errorf("format: '%v' is invalid", ps.Format)
				}); xv == nil {
					fmt.Printf("%v %v Running widget\n", APP_EMOJI, APPLICATION_NAME)
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

	defaultDirectory := "/foo-bar"

	paramSet := assistant.NewParamSet[WidgetParameterSet](widgetCommand)
	paramSet.BindValidatedString(
		assistant.NewFlagInfo("directory", "d", defaultDirectory),
		&paramSet.Native.Directory,
		func(value string) error {
			// ideally, we should check if the Flag has been explicitly set
			//
			if value == defaultDirectory {
				return nil
			}
			if _, err := os.Stat(value); err != nil {
				if os.IsNotExist(err) {
					return err
				}
			}
			return nil
		},
	)

	paramSet.Native.OutputFormatEnumInfo = assistant.NewEnumInfo(assistant.AcceptableEnumValues[OutputFormatEnum]{
		XmlFormatEn:      []string{"xml", "x"},
		JsonFormatEn:     []string{"json", "j"},
		TextFormatEn:     []string{"text", "tx"},
		ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
	})

	paramSet.Native.OutputFormatEn = paramSet.Native.OutputFormatEnumInfo.NewValue()

	paramSet.BindValidatedEnum(
		assistant.NewFlagInfo("format", "f", "xml"),
		&paramSet.Native.OutputFormatEn.Source,
		func(value string) error {
			if paramSet.Native.OutputFormatEnumInfo.En(value) == XmlFormatEn {
				return nil
			}
			return fmt.Errorf("only xml format is currently supported, other formats available in future release")
		},
	)

	paramSet.BindBool(
		assistant.NewFlagInfo("concise", "c", false),
		&paramSet.Native.Concise,
	)

	paramSet.BindValidatedString(
		assistant.NewFlagInfo("pattern", "p", ""),
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
	_ = widgetCommand.MarkFlagRequired("pattern")

	const LO = uint(25)
	const HI = uint(50)
	const DEF = uint(10)

	paramSet.BindValidatedUintWithin(
		assistant.NewFlagInfo("threshold", "t", DEF),
		&paramSet.Native.Threshold,
		LO, HI,
	)

	// If you want to disable the widget command but keep it in the project for reference
	// purposes, then simply comment out the following 2 register calls:
	// (Warning, this may just create dead code and result in lint failure so tread
	// carefully.)
	//
	container.MustRegisterRootedCommand(widgetCommand)
	container.MustRegisterParamSet(WIDGET_PSNAME, paramSet)

	return widgetCommand
}
