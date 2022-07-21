# clingy

A CLI that helps you test other CLIs with end-to-end testing by capturing screenshots of commands in sequence,
so you don't have to.

Currently, this takes screenshots of a window running with imagemagick, and as a result will require X11 running.

Supported platforms:

* Linux

## Requirements

### Running natively

#### Linux

Requires the following dependencies for this to run:

* [imagemagick](https://imagemagick.org/script/download.php)

#### MacOS

Requires the following dependencies to run:

* [Python3](https://www.python.org/downloads/macos/) and [screenshot library](https://pypi.org/project/screenshot/)
* [imagemagick](https://imagemagick.org/script/download.php)

#### Windows

TBD

### Running in docker

This can be run in docker. Instructions TBD

## Misc / Credit

Note - large parts of the organization and structure of this repo were pulled from
[this other repo](https://github.com/aptible/cloud-cli/).
