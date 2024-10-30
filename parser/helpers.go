package parser

import (
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/skelouse/code-to-gpt/prompts"
)

func isIgnoredFile(relativePath string, opts Options) bool {
	filename := filepath.Base(relativePath)

	for _, ignoreFile := range IgnoreFiles {
		if filename == ignoreFile || relativePath == ignoreFile {
			return true
		}
	}

	fileExt := filepath.Ext(filename)
	for _, ignoreExt := range IgnoreExt {
		if fileExt == ignoreExt {
			return true
		}
	}

	// Then, check exclude patterns
	for _, pattern := range opts.Exclude {
		if relativePath == pattern {
			return true
		}

		matched, err := doublestar.Match(pattern, relativePath)
		if err != nil {
			continue
		}
		if matched {
			return true
		}
	}

	// Then, check include patterns
	if len(opts.Include) > 0 {
		for _, pattern := range opts.Include {
			if relativePath == pattern {
				return false
			}

			matched, err := doublestar.Match(pattern, relativePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error matching pattern '%s' with path '%s': %v\n", pattern, relativePath, err)
				continue
			}
			if matched {
				return false
			}
		}
		return true
	}

	// Otherwise, include the file
	return false
}

func isIgnoredDirectory(relativePath string, opts Options) bool {
	dirName := filepath.Base(relativePath)

	// First, check hard ignores
	for _, ignoreDir := range IgnoreDirectories {
		if dirName == ignoreDir || relativePath == ignoreDir {
			return true
		}
	}

	// Then, check exclude patterns
	for _, pattern := range opts.Exclude {
		if relativePath == pattern {
			return true
		}

		matched, err := doublestar.Match(pattern, relativePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error matching pattern '%s' with path '%s': %v\n", pattern, relativePath, err)
			continue
		}
		if matched {
			return true
		}
	}

	// Otherwise, include the directory
	return false
}

func writeFileContents(basePath, path string, newFile *os.File) error {
	relativePath, err := filepath.Rel(basePath, path)
	if err != nil {
		return err
	}

	extension := filepath.Ext(path)
	var (
		commentPrefix string
		commentSuffix string
	)
	switch strings.ToLower(extension) {
	case ".js", ".go", ".java", ".c", ".cpp", ".cs":
		commentPrefix = "//"
	case ".md":
		commentPrefix = "<!--"
		commentSuffix = "-->"
	default:
		commentPrefix = "#"
	}

	_, err = newFile.WriteString(prompts.FileSeparator)
	if err != nil {
		return err
	}

	// Write the comment with the relative path to the "mash.gpt" file
	_, err = newFile.WriteString(fmt.Sprintf("%s %s %s\n", commentPrefix, relativePath, commentSuffix))
	if err != nil {
		return err
	}

	// Open the file
	srcFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Write the contents of the file to the "mash.gpt" file
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

func splitMashFile(mashFile *os.File) error {
	fileInfo, err := mashFile.Stat()
	if err != nil {
		return fmt.Errorf("getting mash file information: %s", err)
	}

	fileSize := fileInfo.Size()
	numFiles := int(math.Ceil(float64(fileSize) / float64(MaxFileLength)))

	buffer := make([]byte, MaxFileLength)
	for i := 1; i <= numFiles; i++ {
		bytesRead, err := mashFile.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("reading mash file: %s", err)
		}

		outputFilename := fmt.Sprintf("%s/send%d-%d.gpt", ToSendDirName, i, numFiles)
		outputFile, err := os.Create(outputFilename)
		if err != nil {
			return fmt.Errorf("creating output file: %s", err)
		}
		defer outputFile.Close()

		_, err = outputFile.WriteString(outputFilename + "\n\n")
		if err != nil {
			return fmt.Errorf("writing output file name: %s", err)
		}

		_, err = outputFile.Write(buffer[:bytesRead])
		if err != nil {
			return fmt.Errorf("writing output file: %s", err)
		}
	}

	return nil
}

func writeMashFile(mashFile *os.File) error {
	outputFilename := fmt.Sprintf("%s/send.gpt", ToSendDirName)
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("creating output file: %s", err)
	}
	defer outputFile.Close()

	data, err := io.ReadAll(mashFile)
	if err != nil {
		return fmt.Errorf("reading mashFile failed with: %s", err)
	}
	_, err = outputFile.Write(data)
	if err != nil {
		return fmt.Errorf("writing to outputFile(%s) failed with: %s", outputFilename, err)
	}

	return nil
}
