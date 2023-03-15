## Installation

```shell
go install github.com/skelouse/code-to-gpt@latest
```

## Usage

First create your prompt which will outline the expectations of chatGPT.  Simply modify the example prompt below to fit your needs, and paste the content of the `send.gpt` file where indicated.


```shell
# from inside a project you would like to parse

code-to-gpt

# then copy the contents of `send.gpt` and paste it below the prompt
```

### Example prompt

```
// Quick summary of the code

I have a Go program that reads the contents of the current working directory and its subdirectories up to a maximum recursion depth of 20. The program writes the relative path of each file as a comment, followed by the file's content, to a file named "send.gpt".

// Explain what is about to happen

I will paste the contents of the generated "send.gpt" file below. Please review the code and provide any suggestions, improvements, or modifications.

// Fill out check boxes for what you would like done, and add a new one if it is not listed.

[ ] Refactor the code
[ ] Summarize the code
[X] Optimize performance
[ ] Identify potential bugs
[ ] Suggest better naming conventions
[ ] Provide additional error handling
[ ] Suggest additional features

Here is the content of the "send.gpt" file:

<content of send.gpt>
```





