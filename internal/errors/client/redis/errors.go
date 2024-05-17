package redis

import (
	"fmt"
	"strings"
)

// prefixFmt is the format for the error prefix
const prefixFmt = "redis error: %v"

// WrapError wraps the given error with the redis error prefix
func WrapError(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(prefixFmt, err)
}

// IsWrappedError checks if the given error is wrapped with the redis error prefix
func IsWrappedError(err error) bool {
	return err != nil && strings.Index(err.Error(), prefixFmt) >= 0
}
