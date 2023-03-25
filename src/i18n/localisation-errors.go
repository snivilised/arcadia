package i18n

import (
	"fmt"

	"golang.org/x/text/language"
)

// ❌ FailedToCopyOptions

// FailedToCopyOptions options copy failed
func FailedToCopyOptionsNativeError(tag language.Tag) error {
	return fmt.Errorf(
		"i18n: failed to copy options whilst using language '%v'", tag,
	)
}
