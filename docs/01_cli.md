# CLI Reference

The CLI is powered by [Cobra](https://github.com/spf13/cobra), which has some pretty standard
features/naming conventions.

## Common Flags

* `--help` - on any command, you can view additional information with this flag
* `-i` or `--inputFile` - commands that take an `inputFile` will use this as a clingy YAML template. If
  none specified, will default to `.clingy.yaml`.
* `-o` or `--outputPath` - on any command if an output path is provided, use that instead. defaults to `./output`
* `-r` or `--reportStyle` - on any running or report-based command, this changes the desired [report output format](/03_outputs)
* `-u` or `--unixTimestampDirDisabled` - this will disable unix timestamp subdirectory-based work and instead save directly
to the specified to the location from `--outputPath` (or `./output` if unspecified)

{==

For debugging, you can also run in debug mode with `-d` or `--debug` and additional more verbose logs will be placed
in the `--outputPath` with details.

==}

## Commands

### Clean

`clean` will empty the target output directory, usually specified by `--outputPath`, so you can repeatedly
run with a clean destination directory, especially if you are leaving unix timestamp subdirectory-based work.

{==

__Warning__ - using this command will clean out entire output directory. Make sure it's safe to purge!

==}

### Completion

`completion` will install autocompletion script for the current shell. Provided by 
[Cobra](https://github.com/spf13/cobra).

### Help

`help` prints out the same help text as `--help`. Provided by [Cobra](https://github.com/spf13/cobra).

### Init

`init` will set up a sample `.clingy.yaml` in the current directory. The default file is generated from
[this snippet](https://github.com/madhuravius/clingy/blob/main/cmd/init.go#L11).

This can be overridden with a desired path at `-i <PATH YOU SPECIFY`.

### Run

`run` will run a specified execution from a clingy YAML file specified by the `--inputFile` (or `.clingy.yaml` by default).

This function defaults to a [simple HTML report](/03_outputs/#html-reports-simple). You can change this with
the `--reportStyle` flag.

### Validate

`validate` will check to see if you have an environment that will support running clingy. See [installation
and usage instructions](/#installation) for details. It will also validate the `--inputFile` passed file or
the default `.clingy.yaml`. Run this before running clingy to check for errors before trying to execute it.

### Version

`version`, run this to print out the current version of clingy.