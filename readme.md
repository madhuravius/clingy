# clingy

A CLI that helps you test other CLIs with end-to-end testing by capturing screenshots of commands in sequence,
so you don't have to.

## TODO

- [x] Take an input yaml and execute a set of commands in order
- [x] For each command execution, capture a screenshot of its output
- [x] Validate colorized support
- [x] Add basic label/description of each image
- [ ] Use [go bindings for imagemagick](https://github.com/gographics/imagick) instead of cmd exec
- [ ] Add examples
- [ ] Add multiline displaying text
- [ ] Add options to how Label/Description are displayed including placing text outside the body of an image
- [ ] Add docker example/support and validate
- [ ] Output clingy results to HTML with image/video references for easy reuse
- [ ] Add support for user interaction (keystroke entry) instead of only commands
- [ ] Waiters for waiting for an action to finish or anticipated text in payload
- [ ] Recordings

## Requirements

### Running natively

Requires the following dependencies for this to even run:

* [imagemagick](https://imagemagick.org/script/download.php)

### Running in docker

This can be run in docker. Instructions TBD

## Misc / Credit

Note - large parts of the organization and structure of this repo were pulled from
[this other repo](https://github.com/aptible/cloud-cli/).
