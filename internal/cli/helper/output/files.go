package output

import (
	"os"
)

// MakeFile creates a file with the specified path and content
func MakeFile(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)

	// If the file or directory doesn't exist, try to create it
	if err != nil && os.IsNotExist(err) {
		dir := parseDirPath(path)
		err = MakeDir(dir)
		if err != nil {
			return err
		}

		return os.WriteFile(path, content, 0644)
	}

	return err
}

// MakeDir creates a directory with the specified path
func MakeDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func parseDirPath(path string) string {
	filename := parseFileName(path)

	return path[:len(path)-len(filename)]
}

func parseFileName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}

	return path
}
