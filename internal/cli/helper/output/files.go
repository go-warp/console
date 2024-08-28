package output

import "os"

// MakeFile creates a file with the specified path and content
func MakeFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}
