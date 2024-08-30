package cli

import "os/exec"

// FixGoimports fixes the imports in the file
func FixGoimports(path string) error {
	return exec.Command("goimports", "-w", path).Run()
}
