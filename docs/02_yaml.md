# clingy YAML file

By default, `clingy init` will place a `.clingy.yaml` file in your present working directory.

## Top-level Parameters

* `label` - a label to label the entire clingy workflow
* `description` - provide a description for the clingy workflow

## Steps

### Command

Specify a command to execute (ex: `echo`). 

```yaml
steps:
  - command: <INSERT A COMMAND HERE>
    ...
```

### Arguments (and Inputs)

In an array of `args`, provide a set of arguments to execute against. 

```yaml
steps:
  - # ... other step fields
    args:
      - <INSERT AN ARGUMENT HERE>
      - <INSERT ANOTHER ARGUMENT HERE>
      # ... keep adding arguments
```

If you want to use an input set by an output, you can use it with `$[[KEY NAME]]`, where `KEY NAME`
is a value you specify in a `key` in a preceding step. This is highlighted in the [outputs section 
below](/02_yaml/#outputs).

### Outputs

There are a few different ways you can capture outputs, and then, in turn, use them as inputs. For example,
say you were to use `#!shell echo "Hello my name is Madhu!"`, you may want to extract the name from that echo,
or echo it right back up, with the intended use of:  `#!shell echo "Hello $[[name]]!"` resulting in
`Hello Madhu!`. This document will outline the major ways in which this is possible:

* __Full Output__ - capture the entirety of the output and store it for reuse
* __Regex__ - evaluate a regular expression and use the first match it encounters
* __Positional__ - target a delimiter (ex: `|`) and specify its index

There are also some arguments and helpers that transcend all output extraction mechanics. See the

The overarching structure is that any output of a command you wish to store will need to be stored in
an `output_processing` field:

```yaml
steps:
  - # ... other step fields
    output_processing:
      key: <INSERT A KEY HERE>                                # a key to search on later for reuse
      matching_type: <CHOOSE ONE OF THE MATCHING TYPES BELOW> # a matching type for use later
      # ... other fields specific to matching type, depending on which you chose
  - # ... other step fields
    command: <INSERT A COMMAND HERE>
    args:
      # use the key you specified above under "key"
      - $[[KEY NAME SPECIFIED ABOVE HERE FROM <INSERT A KEY HERE>]]
```

You can also see this in the Go code [here](https://github.com/madhuravius/clingy/blob/main/lib/models.go#L43),
for the `ClingyOutputProcessing` struct.

#### Full Output

This is used when `matching_type` is set to `full`.

This will fully use the output of a command and store it in a map with a corresponding `key` for reuse.

##### Example

```yaml
label: "..."
steps:
  - label: saving input
    command: echo
    args:
      - sample-input-value
    output_processing:
      key: sample_input_key
      matching_type: full
  - label: using output
    command: echo
    args:
      - $[[sample_input_key]]
```

This can also be found in this test case: [02_basic_full_input_will_pass.yaml](https://github.com/madhuravius/clingy/blob/main/cmd/test_data/02_basic_full_input_will_pass.yaml).

If done properly, the effect of this would be but in staggered screenshots:

```sh
# step 1 runs
> echo sample-input-value
sample-input-value
# step 2 runs
> echo $[[sample_input_key]]
sample-input-value
```

#### Regex

This is used when `matching_type` is set to `regex`.

This matching type will evaluate an additional `regex` field within `maching_args`, evaluating 
a regex you provide. If found, it will assign the corresponding first match to the `key` for reuse.

##### Example

```yaml
label: "..."
steps:
  - label: saving input
    command: echo
    args:
      - "[[sample]]-input-value"
    output_processing:
      key: sample_input_key
      matching_args:
        regex: \[\[([^]]+)\]
      matching_type: regex
  - label: using output
    command: echo
    args:
      - $[[sample_input_key]]
```

This can also be found in this test case: [04_basic_regex_input_will_pass.yaml](https://github.com/madhuravius/clingy/blob/main/cmd/test_data/04_basic_regex_input_will_pass.yaml).

If done properly, the effect of this would be but in staggered screenshots:

```bash
# step 1 runs
> echo "[[sample]]-input-value"
[[sample]]-input-value
# step 2 runs, pulling the value out of [[ ]]
> echo $[[sample_input_key]]
sample
```

#### Positional

This is used when `matching_type` is set to `positional`.

This matching type will evaluate:
* a `positional_delimiter` field within `matching_args`, splitting a string by that delimiter
* a `positional_index` field that requires an integer you select  from the delimited string above

A combination of the two will result in a value, which is then stored in a map and can be retrieved
by `key`.

##### Example

```yaml
label: "..."
steps:
  - label: saving input
    command: echo
    args:
      - sample-input-value
    output_processing:
      key: sample_input_key
      matching_args:
        positional_delimiter: "-"
        positional_index: 0  # this will result in "sample" from above, if you used "1" it would be "input"
      matching_type: positional
  - label: using output
    command: echo
    args:
      - $[[sample_input_key]]
```


This can also be found in this test case: [06_basic_positional_input_will_pass.yaml](https://github.com/madhuravius/clingy/blob/main/cmd/test_data/06_basic_positional_input_will_pass.yaml).

If done properly, the effect of this would be but in staggered screenshots:

```sh
# step 1 runs
> echo "sample-input-value"
sample-input-value
# step 2 runs, pulling the value out of [sample, input, key] of index 0, or "sample"
> echo $[[sample_input_key]]
sample
```

#### Other Arguments

##### Fail on no matches

If you are anticipating these and should fail, you can use the `fail_on_no_match` to `true`. This will
raise an error and interrupt execution of the suite. For example:

```yaml
label: "..."
steps:
  - label: saving input
    command: echo
    args:
      - hello
    output_processing:
      key: sample_input_key
      matching_type: regex
      regex: ^bye$             # this will fail, but ^hello$ will pass!
      fail_on_no_match: true
```

This will error because the regex did not find that echo statement and the suite will also exit with
an error code 1.