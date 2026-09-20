package fileutils

import (
	"bufio"
	"os"
)

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
