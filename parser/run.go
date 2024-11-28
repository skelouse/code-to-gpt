package parser

import (
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
)

type Options struct {
	SplitFiles bool
	Clipboard  bool
	Include    []string
	Exclude    []string
}

func Run(opts Options) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current working directory: %s", err)
	}

	// Remove the send directory if it exists
	if _, err := os.Stat(ToSendDirName); err == nil {
		os.RemoveAll(ToSendDirName)
	}

	mashFile, err := os.Create(MashFileName)
	if err != nil {
		return fmt.Errorf("creating mash file: %s", err)
	}

	// Clean up the mash file when we're done
	defer func() {
		mashFile.Close()
		err := os.Remove(MashFileName)
		if err != nil {
			panic(fmt.Errorf("removing mash file: %s", err))
		}
	}()

	filesProcessed := new([]string)
	err = parse(ParseOptions{
		baseOptions:         opts,
		basePath:            cwd,
		currentPath:         cwd,
		newFile:             mashFile,
		remainingRecursions: MaxRecursions,
		filesProcessed:      filesProcessed,
	})
	if err != nil {
		return fmt.Errorf("parsing files: %s", err)
	}

	mashFile.Seek(0, 0)
	if opts.Clipboard {
		if opts.SplitFiles {
			return fmt.Errorf("cannot copy multiple files to clipboard")
		}

		data, err := io.ReadAll(mashFile)
		if err != nil {
			return fmt.Errorf("splitting mash file: %s", err)
		}

		err = clipboard.WriteAll(string(data))
		if err != nil {
			return fmt.Errorf("writing to clipboard: %s", err)
		}

	} else { // Write to sendGPT directory
		// Create the file directory
		if err := os.MkdirAll(ToSendDirName, os.ModePerm); err != nil {
			return fmt.Errorf("creating send directory: %s", err)
		}

		if opts.SplitFiles {
			// Split the mash file into smaller files to be consumed by chat bot
			err = splitMashFile(mashFile)
			if err != nil {
				return fmt.Errorf("splitting mash file: %s", err)
			}
		} else {
			err = writeMashFile(mashFile)
			if err != nil {
				return err
			}

		}
	}

	// Build the tree and print it
	root := buildTree(*filesProcessed)
	printTree(root, "", true)

	// Print counts
	_, fileCount := countNodes(root)

	message := ""
	if opts.Clipboard {
		message = "copied to clipboard"
	} else {
		message = "written to ./sendGPT/"
	}

	fmt.Printf("\n%d files %s\n", fileCount, message)

	return nil
}
