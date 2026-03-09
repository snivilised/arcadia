package locale_test

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const (
	GrafficoSourceID = "github.com/snivilised/graffico"
)

// GrafficoData provides the source identifier used by the unit-test
// messages for the graffico project.
type GrafficoData struct{}

func (td GrafficoData) SourceID() string {
	return GrafficoSourceID
}

// 🧊 Pavement Graffiti Report

// PavementGraffitiReportTemplData contains template data for reporting
// that graffiti has been found on a pavement, including its primary
// colour.
type PavementGraffitiReportTemplData struct {
	GrafficoData
	Primary string
}

func (td PavementGraffitiReportTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "pavement-graffiti-report.graffico.unit-test",
		Description: "Report of graffiti found on a pavement",
		Other:       "Found graffiti on pavement; primary colour: '{{.Primary}}'",
	}
}

// ☢️ Wrong Source Id

// WrongSourceIDTemplData contains template data that deliberately uses
// an incorrect source identifier to verify error handling in tests.
type WrongSourceIDTemplData struct {
	GrafficoData
}

func (td WrongSourceIDTemplData) SourceID() string {
	return "FOO-BAR"
}

func (td WrongSourceIDTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "wrong-source-id.graffico.unit-test",
		Description: "Incorrect Source Id for which doesn't match the one n the localizer",
		Other:       "Message with wrong id",
	}
}
