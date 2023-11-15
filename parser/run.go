package parser

import (
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
)

type Options struct {
	WithPrompt bool
	SplitFiles bool
	Clipboard  bool
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

	err = parse(cwd, cwd, mashFile, MaxRecursions)
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

		return clipboard.WriteAll(string(data))
	}

	// Create the file directory
	if err := os.MkdirAll(ToSendDirName, os.ModePerm); err != nil {
		return fmt.Errorf("creating send directory: %s", err)
	}

	if opts.SplitFiles {
		// Split the mash file into smaller files to be consumed by chat bot
		err = splitMashFile(mashFile, opts.WithPrompt)
		if err != nil {
			return fmt.Errorf("splitting mash file: %s", err)
		}
	} else {
		err = writeMashFile(mashFile, opts.WithPrompt)
		if err != nil {
			return err
		}

	}

	return nil
}
