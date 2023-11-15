package parser

var (
	MaxRecursions     = 20
	MaxFileLength     = 10000
	ToSendDirName     = "sendGPT"
	MashFileName      = "mash.gpt"
	IgnoreFiles       = []string{MashFileName}
	IgnoreExt         = []string{".png", ".jpg", ".gpt"}
	IgnoreDirectories = []string{".git", ToSendDirName}
)
