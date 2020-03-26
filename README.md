## Questionator

Convert questions from a CSV file to a JSON file to be read by the Feud game engine.

### How To Build

On a system with GNU Make and Golang installed, change to the project directory and simply run `make` or `make build`, they do the same thing. If you don't have Make installed, you can run `go build`.

### How To Use

Questionator expects questions be provided with a specific type of CSV field structure:

```
question, answer one text, answer one value, answer two text, answer two value
```

The quantity of answers is not limited, but the question must be the first field on the row and the answers and their values must alternate. Each question can have a variable amount of answers and values - they do not need to be the same quantity. Sample CSV files can be located in the project directory.

Questionator has three different run time flags available during execution:

* `input` / `i` - used to specify the input CSV file, defaults to `questions.csv`
* `output` / `o` - used to specify the output JSON file, defaults to `gamedata.json`
* `tabs` / `t` - used to notate the source CSV is tab-delimited, defaults to `false` and Questionator assumes the CSV is comma-delimited

Questionator prints output onto `stdout` in addition to writing the output file.

#### Examples

```shell script
./questionator -i samplequestions-tab.csv -o mygamefile.json -t
```

```shell script
./questionator -i samplequestions-comma.csv
```