package locale

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// UsingConfigFileTemplData holds template data for the message
// that reports which configuration file is being used.
type UsingConfigFileTemplData struct {
	arcadiaTemplData
	// ConfigFileName represents the configuration file name.
	ConfigFileName string
}

func (td UsingConfigFileTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "using-config-file",
		Description: "Message to indicate which config is being used",
		Other:       "Using config file: {{.ConfigFileName}}",
	}
}
