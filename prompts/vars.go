package prompts

import (
	_ "embed"
	"strings"
)

var (
	//go:embed initial.md
	InitialPrompt          string
	LastFileSuffix         = " (last-file)\n\n"
	IntermediateFileSuffix = ` (please, don't forget the original prompt, respond with "learned".)` + "\n\n"
	FileSeparator          = strings.Repeat("-", 79) + "\n"
)
