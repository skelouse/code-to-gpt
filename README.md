# code-to-gpt

`code-to-gpt` is a Go program that reads the contents of the current working directory and its subdirectories up to a maximum recursion depth of 20. The program writes the relative path of each file as a comment, followed by the file's content, to multiple files named "send1-3.gpt", "send2-3.gpt", and "send3-3.gpt" (representing 1 out of 3, 2 out of 3, and 3 out of 3, respectively). This tool is useful for analyzing and discussing code projects with AI language models like OpenAI's ChatGPT.

## Simple `code-to-gpt -c` example

[code-to-gpt-ex.webm](https://github.com/skelouse/code-to-gpt/assets/42463301/842eca5d-f6d4-4dd3-b89f-c8b6503cb7bb)

## Installation

```shell
go install github.com/skelouse/code-to-gpt@latest
```

## Usage

1. Run `code-to-gpt`
2. Paste into chat.
3. ???
4. Write great code.

### Example Session

```shell
# Command with new flags
code-to-gpt --split-files

# Or if you want to copy the output to the clipboard
code-to-gpt --c

```


## Super Diff

The super-diff.sh script is a utility to output all un-staged changes to a git repository. It works by performing a git diff on un-staged files, displaying the output in the terminal. You can easily redirect the output to a file or copy it to your clipboard.

### Why?

When working with ChatGPT and running super-diff.sh, you can generate a clean representation of your changes and incorporate it into a prompt. This way, you can maintain the context of your work and get more accurate responses.  I have personally used this when developing an API, simply prompt chatGPT for the initial API calls, then use super-diff.sh to show what was done.  Then I can paste it back in as `Awesome thanks! This is what I did {super-diff.sh output} now I would like to implement calls X, Y and Z`.

Note: If you have a problem `git reset`

### Examples

```shell
# Redirect output to a file

$ ./super-diff.sh > diff.txt
```

```shell
# Copy output to clipboard (Linux)
$ ./super-diff.sh | xclip -selection clipboard

# Copy output to clipboard (MacOS)
$ ./super-diff.sh | pbcopy
```

You can also write it as a function in your bashrc or zshrc file:

```shell
# (linux)
function super-diff() {
  ./super-diff.sh | xclip -selection clipboard
}

# (MacOS)
function super-diff() {
  ./super-diff.sh | pbcopy
}
```
