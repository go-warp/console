package init

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	boilerplaceRepoUrl = "https://github.com/go-warp/warp-boilerplate/archive/refs/heads/master.zip"
	tmpZipFile         = "/tmp/sitnikovik-warp-boilerplate.zip"
	defaultDestination = "."
)

// NewCommand creates a command that is used to configure the Warp
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init the warp framework project",
		RunE:  run,
	}
	cmd.Flags().String("dest", defaultDestination, "Destination path to extract the project boilerplate")

	return cmd
}

// run runs the command
func run(cmd *cobra.Command, args []string) error {
	conf := parseConfig(cmd)

	// Downloading the boilerplate archive
	if err := downloadFile(tmpZipFile, boilerplaceRepoUrl); err != nil {
		return err
	}

	// Extracting zip
	if err := unzip(tmpZipFile, conf.destination); err != nil {
		return fmt.Errorf("err while extraction boilerplate: %w", err)
	}
	os.Remove(tmpZipFile)

	return nil
}

// downloadFile загружает файл по URL и сохраняет его на диск
func downloadFile(filepath string, url string) error {
	// Create the file to download in
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Downloading
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Copying
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// unzip extracts the ZIP archive to the specified destination
func unzip(src string, dest string) error {
	// Open the ZIP archive
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		// Construct the full file path
		fpath := filepath.Join(dest, file.Name)

		// Check if the file is a directory
		if file.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			// Create parent directories if they don't exist
			os.MkdirAll(filepath.Dir(fpath), os.ModePerm)

			// Create the file
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			// Open the file in the ZIP archive for reading
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			// Copy the content from the archive file to the disk file
			_, err = io.Copy(outFile, rc)

			if err != nil {
				return err
			}
		}
	}
	return nil
}
