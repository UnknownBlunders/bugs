# Bugs

A simple go based bug tracker for the cli

## Installation:

Download the latest binary from the [releases page](https://github.com/UnknownBlunders/bugs/releases).

Install it to your preferred location:

``` bash
sudo install bugs-<os>-<release> /usr/local/bin/bugs
```
Example:
``` bash
sudo install bugs-ubuntu-v0.0.3 /usr/local/bin/bugs
```

## Usage:

Bugs will track bugs in `.buglist.json`. If there is not a `.buglist.json` in the directory that you ran bugs in, then bugs will create one. This is perfect for keeping a buglist in each of your git repos!

The help command will show you full usage:

``` bash
bugs help
```

