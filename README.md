rgen
==========

Easy to use helpers for random data generation

[![push](https://github.com/n0str/rgen/workflows/push/badge.svg?branch=master&event=push)](https://github.com/n0str/rgen/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/n0str/rgen?)](https://goreportcard.com/report/github.com/n0str/rgen)

Usage
-----

Run binary file with the following arguments

```
Usage: ./rgen: <command> <arguments>
Commands: id58, id58l
Arguments:
  -n int
        how many symbols to generate (default 10)
```

Output
-----

Sample output of `for i in {1..5}; do ./rgen id58 -n 16; done`
```
qGGXyV3JfbAsFFDy
gFDvKc2uNbeCnuMj
8vBJxASnbdPXFFeX
8ssaJNpYkqUep9tB
MoCQpSPTwouywycE
```

Sample output of `for i in {1..5}; do ./rgen id58l -n 4; done`
```
pa4u
sc99
ua2d
eohp
o3wr
```

Installation
-----

```
go get github.com/n0str/rgen
```

You can also download release files from [GitHub](https://github.com/n0str/rgen/releases/tag/v1.1)

Build
-----

Clone the repository, compile binary for linux or macos and run.

```
git clone https://github.com/n0str/rgen
cd rgen
GOOS=linux go build -o rgen-linux .
```

Now you can run the binary file `rgen-linux`