package translate

import (
	"golang.org/x/text/language"

	"github.com/snivilised/arcadia/src/internal/l10n"
)

func GetLanguageNotSupportedErrorMessage(tag language.Tag) string {
	data := l10n.LanguageNotSupportedTemplData{
		Language: tag.String(),
	}

	return localise(data)
}

func GetCouldNotLoadTranslationFileErrorMessage(tag language.Tag, fullPath string) string {
	data := l10n.CouldNotLoadTranslationFileTemplData{
		Language: tag.String(),
		FullPath: fullPath,
	}

	return localise(data)
}
