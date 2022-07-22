# clingy

[![GitHub release](https://img.shields.io/github/release/madhuravius/clingy)](https://github.com/madhuravius/clingy/)
[![Test](https://github.com/madhuravius/clingy/actions/workflows/test.yaml/badge.svg)](https://github.com/madhuravius/clingy/actions/workflows/test.yaml)
[![license](https://img.shields.io/github/license/madhuravius/clingy.svg)](https://github.com/madhuravius/clingy/blob/main/LICENSE.md)

A CLI that helps you test other CLIs with end-to-end testing by capturing screenshots of commands in sequence,
so you don't have to.

See docs here: [Link](https://madhuravius.github.io/clingy/)

Supported platforms:

* Linux
* Mac

## Requirements and Installation

For installation instructions, please see this doc for guidance: [Link](https://madhuravius.github.io/clingy/)

### Running natively

#### Linux

Requires the following dependencies for this to run:

* [imagemagick](https://imagemagick.org/script/download.php)

#### MacOS

Requires the following dependencies to run:

* [Python3](https://www.python.org/downloads/macos/) and [screenshot library](https://pypi.org/project/screenshot/)
* [imagemagick](https://imagemagick.org/script/download.php)

These must both be present in your path to function.

#### Windows

TBD

### Running in docker

This can be run in docker. Instructions TBD

## Misc / Credit

Credit - I was inspired from the Makefile and initial structuring of
[this other repo](https://github.com/aptible/cloud-cli/) when propping up this repository.
