## file-split ##

A tool to help extract a portion of a file

### Usage ###

Currently the only functionality is to take one contiguous portion of a file and either write it to stdout or a file.

For example, if I wanted to output bytes 100 to 399 of file.dat to portion.dat I would do the following:
```bash
> file-split -file=file.dat -output=portion.dat -start=100 -end=400
```
Defaults for `-start` are the beginning of the file and `-end` the end of the file, so if I wanted the first 100 bytes
of a file printed to stdout, I would do the following:
```bash
> file-split -file=file.dat -end=100
```
