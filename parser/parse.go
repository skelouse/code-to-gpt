package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func parse(basePath, currentPath string, newFile *os.File, remainingRecursions int) error {
	// Return if the recursion limit has been reached
	if remainingRecursions == 0 {
		return nil
	}

	// Read the contents of the current directory
	files, err := ioutil.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(currentPath, file.Name())

		if file.IsDir() {
			if !isIgnoredDirectory(file.Name()) {
				err := parse(basePath, path, newFile, remainingRecursions-1)
				if err != nil {
					return err
				}
			}
		} else {
			if !isIgnoredFile(file.Name()) {
				err = writeFileContents(basePath, path, newFile)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
