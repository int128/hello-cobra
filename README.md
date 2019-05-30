# hello-cobra

Hello world with [spf13/cobra](https://github.com/spf13/cobra).

```
% ./hello-cobra --help
Hello world with cobra.

Examples:
  # Root
  ./hello-cobra

  # Echo
  ./hello-cobra echo

Usage:
  ./hello-cobra [flags]
  ./hello-cobra [command]

Available Commands:
  echo        Just say args
  help        Help about any command

Flags:
      --foo string   foo (root)
  -h, --help         help for ./hello-cobra

Use "./hello-cobra [command] --help" for more information about a command.

% ./hello-cobra
2019/05/30 11:25:42 hello
```

```
% ./hello-cobra echo --help
Just say args

Usage:
  ./hello-cobra echo [flags]

Flags:
      --bar string   bar (echo)
  -h, --help         help for echo

% ./hello-cobra echo foo
2019/05/30 11:25:30 args=[foo]
```
