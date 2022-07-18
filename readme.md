# clingy

A CLI that helps you test other CLIs with end-to-end testing by capturing screenshots of commands in sequence,
so you don't have to.

Currently, this takes screenshots of a window running with imagemagick, and as a result will require X11 running.

## Requirements

### Running natively

Requires the following dependencies for this to even run:

* [imagemagick](https://imagemagick.org/script/download.php)

### Running in docker

This can be run in docker. Instructions TBD

## Misc / Credit

Note - large parts of the organization and structure of this repo were pulled from
[this other repo](https://github.com/aptible/cloud-cli/).
