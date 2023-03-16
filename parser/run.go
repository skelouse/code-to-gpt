package parser

import (
	"fmt"
	"os"
)

func Run() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current working directory: %s", err)
	}

	// Remove the send directory if it exists
	if _, err := os.Stat(ToSendDirName); err == nil {
		os.RemoveAll(ToSendDirName)
	}

	if err := os.MkdirAll(ToSendDirName, os.ModePerm); err != nil {
		return fmt.Errorf("creating send directory: %s", err)
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

	// Split the mash file into smaller files to be consumed by chat bot
	mashFile.Seek(0, 0)
	err = splitMashFile(mashFile)
	if err != nil {
		return fmt.Errorf("splitting mash file: %s", err)
	}

	return nil
}
