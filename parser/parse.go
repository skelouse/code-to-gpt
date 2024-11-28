package parser

import (
	"fmt"
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
	currentSize         int64
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
					currentSize:         opts.currentSize,
				})
				if err != nil {
					return err
				}
			}
		} else {
			if !isIgnoredFile(relativePath, opts.baseOptions) {
				// Get file size
				fileInfo, err := os.Stat(path)
				if err != nil {
					return err
				}

				// Calculate new total size
				fileSize := fileInfo.Size()
				newTotalSize := opts.currentSize + fileSize

				// Check if new total size exceeds MaxSize
				if opts.baseOptions.MaxSize > 0 && newTotalSize > opts.baseOptions.MaxSize {
					return fmt.Errorf("\n...maximum size limit reached at %s (%d bytes).\n...current limit (%d bytes), maybe set --max-size to a value higher than that", path, fileSize, opts.baseOptions.MaxSize)
				}

				// Write file contents
				err = writeFileContents(opts.basePath, path, opts.newFile)
				if err != nil {
					return err
				}

				// Update current size
				opts.currentSize = newTotalSize

				*opts.filesProcessed = append(*opts.filesProcessed, relativePath)
			}
		}
	}

	return nil
}
