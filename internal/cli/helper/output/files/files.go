package files

import (
	"os"
	"os/exec"
)

// Create creates a file with the specified path and content
func Create(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)

	// If the file or directory doesn't exist, try to create it
	if err != nil && os.IsNotExist(err) {
		dir := parseDirPath(path)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

		return os.WriteFile(path, content, 0644)
	}

	return err
}

// FixGoimports fixes the imports in the file
func FixGoimports(path string) error {
	return exec.Command("goimports", "-w", path).Run()
}

// parseDirPath returns the directory path of the specified file path
func parseDirPath(path string) string {
	filename := parseFileName(path)

	return path[:len(path)-len(filename)]
}

// parseFileName returns the file name of the specified file path
func parseFileName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}

	return path
}
