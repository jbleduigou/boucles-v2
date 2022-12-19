<p align="center"><img width="50%" src="./christmas.svg" alt="Nedeleg Laouen"></p>

# Boucles V2

## Introduction

Boucles v2 is a utility that extracts a list of timecodes from a crossover in xslx format.  
The output file in CSV format is then usable in Nuendo.  
This is a follow up on the [boucles v1](https://github.com/jbleduigou/boucles) project, which had a pdf as the input file.

This v2 version is implemented in [Go](https://go.dev/).

## Basic Usage

### Prerequisites

You will need to have [Go](https://go.dev/) installed on your computer first.  
The easiest way is to use [homebrew](https://brew.sh/) for that:
```bash
brew install golang
```

Also, you need to have [Make](https://www.gnu.org/software/make/) installed on your computer.  
Again, you can use [homebrew](https://brew.sh/) for that.
```bash
brew install make
```

### Installing boucles-v2

Then you can simply use the Makefile to install the utility:
```bash
make install
```

### Converting all files in directory

Once the utility is installed you can run it to convert all files in directory.  
The first argument will the directory containing the files.
```bash
boucles-v2 .
```
The output files will be written to a file with the same name but a different suffix.  
For instance if the input file is `my-movie.xlsx` then the output will be `my-movie.csv`.

The content of the output file will be similar to:
```csv
1,01:00:45:01,01:01:06:15
2,01:01:06:15,01:02:01:03
3,01:02:01:03,01:02:46:10
```

## Contributing

Contributions are welcome!  
Open a pull request to fix a bug, or open an issue to discuss a new feature or change.

## Licenses

This program is under the terms of the BSD 3-Clause License.  
See [https://opensource.org/licenses/BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause).