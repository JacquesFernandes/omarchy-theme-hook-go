package fileutils

import (
	"bufio"
	"os"
	"path/filepath"
)

const TEMP_DIR_BASE = "omarchy-theme-hook-go"

// CopyFile is a util function to do a buffered-copy of a text file from sourcePath to destPath
func CopyFile(sourcePath, destPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	sourceScanner := bufio.NewScanner(sourceFile)
	destWriter := bufio.NewWriter(destFile)

	for sourceScanner.Scan() {
		line := sourceScanner.Text()

		if _, err := destWriter.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err = sourceScanner.Err(); err != nil {
		return err
	}

	if err = destWriter.Flush(); err != nil {
		return err
	}

	return nil
}

// CreateTempProjectFile creates a temp-file using `filename` and appending it with a random suffix.
// This file is created in the directory OS_TEMP_DIR/omarchy-theme-hook-go/`dirName`
// The resulting file has it's file handler returned, if successful.
func CreateTempProjectFile(dirName, filename string) (*os.File, error) {
	tempDirPath := os.TempDir()
	err := os.MkdirAll(filepath.Join(tempDirPath, TEMP_DIR_BASE, dirName), 0o755)
	if err != nil {
		return nil, err
	}

	return os.CreateTemp(tempDirPath, filename)
}
