# code-to-gpt

`code-to-gpt` is a Go program that reads the contents of the current working directory and its subdirectories up to a maximum recursion depth of 20. The program writes the relative path of each file as a comment, followed by the file's content, to multiple files named "send1-3.gpt", "send2-3.gpt", and "send3-3.gpt" (representing 1 out of 3, 2 out of 3, and 3 out of 3, respectively). This tool is useful for analyzing and discussing code projects with AI language models like OpenAI's ChatGPT.

## Installation

```shell
go install github.com/skelouse/code-to-gpt@latest
```

## Usage

1. Change to the directory you would like to analyze
2. Run `code-to-gpt`
3. Locate the generated `send*-*.gpt` files in the `sendGPT` directory
4. Paste the contents of each file into a separate message to chatGPT
5. Once chatGPT has learned the contents of the final file it will provide a list of options for further actions
6. Respond to chatGPT with the number of the option you would like to discuss, or simply respond with what you would like to have done

### Example Session

```
<content of send*-*.gpt sent in separate messages>
```

Response:

```
learned

Based on the contents of the provided files, here are some options for further actions:

Refactor the code
Summarize the code
Optimize performance
Identify potential bugs
Suggest better naming conventions
Provide additional error handling
Suggest additional features
...
```

## TODO

- [ ] Use urfave/cli for parsing flags
- [ ] Automatically generate a prompt based on the code using go-survey or take options from flags
- [ ] Add flag for specifying the prompt directory
- [ ] Add flag for specifying out filename

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
