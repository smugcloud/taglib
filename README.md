# Taglib

Taglib is a helper tool to quickly update metadata for audio files within a specific directory.

## Building

First, clone the repo into the proper location in your $GOPATH:

```
go get -u github.com/smugcloud/taglib
cd $GOPATH/src/github.com/smugcloud/taglib
```

Then, to build:

``` 
make all
```

## Usage

```
$ taglib -h
Audio metadata inspection and updating.

Usage:
  taglib [command]

Available Commands:
  list        List metadata associated with the files in a directory.
  update      Update metadata using the global flags.

Flags:
      --album string       Album name.
      --artist string      Artist name.
      --directory string   Directory to evaluate.

Use "taglib [command] --help" for more information about a command.
```

## Features

- [x] List metadata associated with media files within a directory
- [x] Update Album and Artist metadata 
- [ ] Add additional parameters to update
- [ ] Add recursive updates as an optional parameter 

