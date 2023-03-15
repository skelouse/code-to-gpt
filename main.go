package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	maxRecursions     = 20
	ignoreFiles       = []string{"send.gpt"}
	ignoreDirectories = []string{".git"}
)

func main() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Create a new file
	newFile, err := os.Create("send.gpt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer newFile.Close()

	err = parse(cwd, cwd, newFile, maxRecursions)
	if err != nil {
		fmt.Println("Error parsing directory:", err)
		return
	}
}

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

func isIgnoredFile(filename string) bool {
	for _, ignoreFile := range ignoreFiles {
		if filename == ignoreFile {
			return true
		}
	}
	return false
}

func isIgnoredDirectory(directoryName string) bool {
	for _, ignoreDir := range ignoreDirectories {
		if directoryName == ignoreDir {
			return true
		}
	}
	return false
}

func writeFileContents(basePath, path string, newFile *os.File) error {
	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return err
	}

	extension := filepath.Ext(path)
	comment := ""
	switch strings.ToLower(extension) {
	case ".js", ".go", ".java", ".c", ".cpp", ".cs":
		comment = "//"
	default:
		comment = "#"
	}

	// Write 79 dashes between files
	_, err = newFile.WriteString(strings.Repeat("-", 79) + "\n")
	if err != nil {
		return err
	}

	// Write the comment with the relative path to the "send.gpt" file
	_, err = newFile.WriteString(fmt.Sprintf("%s %s\n", comment, relPath))
	if err != nil {
		return err
	}

	// Open the file
	srcFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Write the contents of the file to the "send.gpt" file
	_, err = io.Copy(newFile, srcFile)
	if err != nil {
		return err
	}

	// Add a newline after the contents
	_, err = newFile.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}
