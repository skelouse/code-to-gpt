package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type ParseOptions struct {
	baseOptions         Options
	basePath            string
	currentPath         string
	newFile             *os.File
	remainingRecursions int
	filesProcessed      *[]string
}

func parse(opts ParseOptions) error {
	// Return if the recursion limit has been reached
	if opts.remainingRecursions == 0 {
		return nil
	}

	// Read the contents of the current directory
	files, err := ioutil.ReadDir(opts.currentPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(opts.currentPath, file.Name())

		relativePath, err := filepath.Rel(opts.basePath, path)
		if err != nil {
			return err
		}

		if file.IsDir() {
			if !isIgnoredDirectory(relativePath, opts.baseOptions) {
				err := parse(ParseOptions{
					baseOptions:         opts.baseOptions,
					basePath:            opts.basePath,
					currentPath:         path,
					newFile:             opts.newFile,
					remainingRecursions: opts.remainingRecursions - 1,
					filesProcessed:      opts.filesProcessed,
				})
				if err != nil {
					return err
				}
			}
		} else {
			if !isIgnoredFile(relativePath, opts.baseOptions) {
				err = writeFileContents(opts.basePath, path, opts.newFile)
				if err != nil {
					return err
				}

				*opts.filesProcessed = append(*opts.filesProcessed, relativePath)
			}
		}
	}

	return nil
}
