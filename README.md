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
Commands: id58, id58l, id58u, id92
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

Sample output of `for i in {1..3}; do ./rgen id58u -n 31; done`
```
2R5TG33TABP9KDLAQ4AY7DZ3EZ4SE85
6NCQNEMNPBMEGW8CBVDKHEQC13V6SJT
JRVM1XC7QTGKSFHH3EALRSYPQQF8MD9
```

Sample output of `for i in {1..3}; do ./rgen id92 -n 9; done`
```
Ke!.rBp2y
R\3s:ynf[
ga8Ue7[w
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