# ctxt

ctxt is a simple CLI tool for categorizing text.

This package aims to categorize the content to be described in the CHANGELOG.

# Install

```
go install github.com/hlts2/ctxt
```

# Usage

```
$ ctxt --help
Categorize text

Usage:
  ctxt [flags]

Flags:
  -f, --file string                 set file path
  -h, --help                        help for ctxt
  -i, --index uint                  set which element of the separation are to be used
  -s, --sep string                  set line separator
      --uncategorized-name string   set uncategorized name (default "others")
  -v, --version                     version for ctxt
```

# Example

The following is an example to integrate with git log command.

```
 git log --pretty=format:"- %s" v0.0.1..master | ctxt --sep=" " --index=1
```
