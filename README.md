[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-cat)](https://repl.it/github/matt4biz/go-class-cat)

# Go class: cat (concatenate files) example
This example shows opening files from command-line arguments and copying their content to standard output.

```shell
$ go run . *.txt
And another one gone,
   and another one gone
Another one bites the dust
```

You should get the same output as calling `cat` with no options.

Exercise: modify this program to handle `-` as a command line argument and read from standard input (i.e., from a pipe).
