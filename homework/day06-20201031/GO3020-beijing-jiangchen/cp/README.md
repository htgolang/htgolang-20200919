# Easy cp command line tool

## How to build

```bash
~# go build
```

## How to use

```bash
~# ./cp --help
cp -s srcFile -d dstFile
  -d string
        destination file path
  -s string
        source file path
```

## Examples

```bash
~# ls
1.txt  README.md  cp  cp.go  go.mod
~# cat 1.txt
this is a test file.
~# ./cp -s 1.txt -d 2.txt
~# ls
1.txt  2.txt  README.md  cp  cp.go  go.mod
~# cat 2.txt
this is a test file.
~# ./cp -s 1.txt -d 2.txt
destination file exists, do you want to overwrite? (y/n)
destination file exists, do you want to overwrite? (y/n)n
abort...
~# ./cp -s 1.txt -d 2.txt
destination file exists, do you want to overwrite? (y/n)y
~# ls
1.txt  2.txt  README.md  cp  cp.go  go.mod
```
