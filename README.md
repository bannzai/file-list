# file-list
file-list is tool for output to file names below passed directory.

## Usage
### simple

```
$ file-list example
example/aaa.txt
example/bbb.txt
example/ccc.txt
example/ddd.txt
example/files.txt
```

### with option

```
$ file-list example --ignore-list example/aaa.txt,example/bbb.txt
example/ccc.txt
example/ddd.txt
example/files.txt
```

Or pass file for wrote ignore output files.
This example ignored `example/aaa.txt`
```
$ file-list example --ignore-with-file example/files.txt
example/bbb.txt
example/ccc.txt
example/ddd.txt
example/files.txt

$ cat example/files.txt
example/aaa.txt
```

### all option

```
$ file-list -h
Options:
        ignore-list                 "Do not output file-list"
        only-list                   "Show only file-list"
        ignore-with-file            "Do not output file-list"
        only-with-file              "Show only file-list"

```

## Install

```
$ go get -u github.com/bannzai/file-list
```

## UseCase

```
$ file-list example | xargs rm
```
