package pg

import (
	"fmt"
	"strings"
)

// pgPrefixFmt is a format string for Postgres errors
const pgPrefixFmt = "postgres error: %v"

// WrapError wraps the given error with the Postgres error prefix
func WrapError(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(pgPrefixFmt, err)
}

// IsWrappedError checks if the given error is wrapped with the Postgres error prefix
func IsWrappedError(err error) bool {
	return err != nil && strings.Index(err.Error(), pgPrefixFmt) >= 0
}
