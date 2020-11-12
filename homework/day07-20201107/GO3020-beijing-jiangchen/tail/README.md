# Easy tail

## How to build

```bash
~# go build
```

## How to use

```bash
~# ./tail --help
  -f string
        file to tail
  -n int
        number of flush line, default is 5 (default 5)
Note: easy version of tail
```

## Examples

```bash
~# ls
README.md  access.log  go.mod  go.sum  main.go  tail  vendor
~# cat access.log
192.168.102.72 - - [28/Jul/2020:15:23:25 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:23:25 +0800] "GET /favicon.ico HTTP/1.1" 404 153 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:23:29 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:23:30 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:23:55 +0800] "GET / HTTP/1.1" 200 616 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:24:02 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:25:15 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:25:48 +0800] "GET /status HTTP/1.1" 200 128966 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:25:48 +0800] "GET /status/format/json HTTP/1.1" 200 1977 "http://192.168.20.248:2020/status" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
192.168.102.72 - - [28/Jul/2020:15:25:49 +0800] "GET /status/format/json HTTP/1.1" 200 2002 "http://192.168.20.248:2020/status" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
~#
~#
~#
~# ./tail -f access.log -n 6
2020/11/12 23:32:37 Seeked access.log - &{Offset:-1071 Whence:2}
line: 192.168.102.72 - - [28/Jul/2020:15:23:55 +0800] "GET / HTTP/1.1" 200 616 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
line: 192.168.102.72 - - [28/Jul/2020:15:24:02 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
line: 192.168.102.72 - - [28/Jul/2020:15:25:15 +0800] "GET /status HTTP/1.1" 501 165 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
line: 192.168.102.72 - - [28/Jul/2020:15:25:48 +0800] "GET /status HTTP/1.1" 200 128966 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
line: 192.168.102.72 - - [28/Jul/2020:15:25:48 +0800] "GET /status/format/json HTTP/1.1" 200 1977 "http://192.168.20.248:2020/status" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
line: 192.168.102.72 - - [28/Jul/2020:15:25:49 +0800] "GET /status/format/json HTTP/1.1" 200 2002 "http://192.168.20.248:2020/status" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0"
```
