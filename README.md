# code-to-gpt

`code-to-gpt -c` parses current working directory and outputs each file to clipboard. Each file will have a comment of the file location.

> For example given a directory:

```
├── lib
│   └── hello.go
```

> Running the following commands

```shell
cd lib
code-to-gpt -c
```

> Will write this to your clipboard

```
-------------------------------------------------------------------------------
// hello.go
package lib

import (
	"fmt"
)

fmt.Println("Hello, world!")
```

## Installation

```shell
go install github.com/skelouse/code-to-gpt@latest
```

## Super Diff

> Note: If you have a problem run `git reset`

The super-diff.sh script is a utility to output all un-staged changes of a local git repository. It uses a diff on un-staged files, displaying the output in the terminal. You can easily redirect the output to a file or copy it to your clipboard.

> Redirect output to a file

```sh
$ ./super-diff.sh > diff.txt
```

> Copy output to clipboard (Linux)

```shell
$ ./super-diff.sh | xclip -selection clipboard
```

> Copy output to clipboard (MacOS)

```shell
$ ./super-diff.sh | pbcopy
```

> Or write it as a function in your .bashrc or .zshrc

> (linux)

```shell
function super-diff() {
  ./super-diff.sh | xclip -selection clipboard
}
```

> (MacOS)

```shell
function super-diff() {
  ./super-diff.sh | pbcopy
}
```
