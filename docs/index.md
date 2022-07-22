# Overview

[![GitHub release](https://img.shields.io/github/release/madhuravius/clingy)](https://github.com/madhuravius/clingy/)
[![Test](https://github.com/madhuravius/clingy/actions/workflows/test.yaml/badge.svg)](https://github.com/madhuravius/clingy/actions/workflows/test.yaml)
[![license](https://img.shields.io/github/license/madhuravius/clingy.svg)](https://github.com/madhuravius/clingy/blob/main/LICENSE)

For source code, see [Github repository](https://github.com/madhuravius/clingy) for details.

This CLI utility helps you automate and test other CLIs. 

## Installation

Releases can be found on this page for download: [Releases Github Page](https://github.com/madhuravius/clingy/releases/)

### Linux

You will need the following installed:

* [imagemagick](https://imagemagick.org/script/download.php) 

You can then download the binary above to proceed by placing it in your `PATH`.

### MacOS

You will need the following installed:

* [imagemagick](https://imagemagick.org/script/download.php)
* [the python library screenshot](https://pypi.org/project/screenshot/)

After these are installed, ensure that `screenshot` has permissions for capturing your screen.
You can then download the binary above to proceed by placing it in your `PATH`.

### Windows

Currently unsupported.

## Basic Usage


{==

__Warning for MacOS users__: For now, you will want to set the environment variable `WINDOW_NAME` 
to whichever application name you are using to execute your commands. For example, if you use 
[iTerm]() or [alacritty](), you will want to set your `export WINDOW_NAME=iTerm` or `export WINDOW_NAME=alacritty` for clingy to know which application to target for screenshots.

==}

See [this link](/clingy/01_yaml/) for details on how to structure the clingy YAML files.

```sh
# MacOS - See above warning and instruction about WINDOW_NAME
# creates a .clingy.yaml file
clingy init

# validate your clingy file and local environment
clingy validate -i ./.clingy.yaml

# runs against the above .clingy.yaml file and save its contents to output
clingy run -i ./.clingy.yaml -o ./output

# you should see an output directory with a new timestamp of an example run that looks like below:
> ls ./output -al
...
drwxr-xr-x  3 user users 4096 Jul 20 00:00 .
drwxr-xr-x 11 user users 4096 Jul 20 00:00 ..
drwxr-xr-x  2 user users 4096 Jul 20 00:00 2158369886 # <---- This contains your report

# clean up artifacts with clean
clingy clean -o ./output
```

## Rationale

I built this tool because I could not find comprehensive CLI/TUI automation tools
out there. I found plenty of tools to automate desktop environments, but not 
one that helped CLI-builders and CLI-testers. I wanted a tool like this to 
help me view and evaluate user-flows similar to what frameworks like 
[Cypress](https://github.com/cypress-io/cypress) or 
[NightwatchJS](https://github.com/nightwatchjs/nightwatch) do for web.

With the emergence of interactive frameworks for TUIs like 
[BubbleTea](https://github.com/charmbracelet/bubbletea) or [Rich](https://github.com/Textualize/rich),
the complexity of interfaces in terminals has skyrocketed.

The goal is to eventually support common actions that CLI/TUI builders and testers
would want to employ in automating common tasks with development. The goals of this tool are as follows:

* __capturing actual usage__ with colors in various dimensions in a terminal emulator
* __input and output__ - be able to pass values from one step to another in a series of commands
* __screenshots/records__ - be able to record common user flows with screenshots or screen recordings
* __annotated reports__ - stitch together recordings/screenshots into a clean browsable report with annotations