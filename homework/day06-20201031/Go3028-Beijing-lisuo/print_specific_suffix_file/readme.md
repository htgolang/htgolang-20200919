Demo
===

## Usage

```bash
root@[10:34:52]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/print_specific_suffix_file]
# ll ../user_manager_proj/
total 40
drwxr-xr-x 6 root root 4096 Nov  2 21:02 ./
drwxr-xr-x 8 root root 4096 Nov  3 09:36 ../
drwxr-xr-x 4 root root 4096 Nov  1 01:08 cmd/
drwxr-xr-x 2 root root 4096 Nov  1 15:07 db/
drwxr-xr-x 2 root root 4096 Nov  1 00:57 define/
-rw-r--r-- 1 root root  179 Nov  2 20:59 go.mod
-rw-r--r-- 1 root root 1266 Nov  1 00:57 go.sum
-rw-r--r-- 1 root root  824 Nov  2 21:02 main.go
-rw-r--r-- 1 root root   94 Nov  2 11:38 README.md
drwxr-xr-x 2 root root 4096 Nov  2 21:01 utils/
root@[10:34:58]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/print_specific_suffix_file]
# go run print_specific_suffix_file.go -f ../user_manager_proj/
file: main.go                     dir: false
root@[10:35:03]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/print_specific_suffix_file]
# ll /opt/
total 74584
drwxr-xr-x  6 root root     4096 Nov  4 10:25 ./
drwxr-xr-x 20 root root     4096 Oct 25 22:03 ../
-rw-r--r--  1 root root       17 Nov  2 22:00 2.go
-rw-r--r--  1 root root        0 Nov  2 21:42 3.cgo
-rw-r--r--  1 root root        0 Nov  2 17:51 check.py
drwx--x--x  4 root root     4096 Aug 25 19:31 containerd/
-rw-r--r--  1 root root   134163 Oct 23 17:51 dex.xml
-rw-r--r--  1 root root   132851 Oct 23 17:51 dex.xmln
lrwxrwxrwx  1 root root       21 Apr 16  2020 go -> ../lib/go-1.13/bin/go
drwxr-xr-x  3 root root     4096 Oct 22 14:49 mongodb-3.2.10/
drwxr-xr-x  3 root root     4096 Oct 22 14:49 mongodb_conf/
-rw-r--r--  1 root root 76066900 Oct 22 14:51 online-3.2.10.zip
drwxr-xr-x  3 root root     4096 Oct 23 17:46 Platform-projs-notes/
-rw-r--r--  1 root root       57 Nov  4 10:25 readme.md
-rw-r--r--  1 root root        0 Nov  2 18:28 src
root@[10:35:07]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/print_specific_suffix_file]
# go run print_specific_suffix_file.go -f /opt/
file: 2.go                        dir: false
file: 3.cgo                       dir: false
root@[10:35:13]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/print_specific_suffix_file]
# go run print_specific_suffix_file.go -f /opt/notafile
2020/11/04 10:35:21 stat /opt/notafile: no such file or directory
exit status 1
root@[10:35:21]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/print_specific_suffix_file]
# go run print_specific_suffix_file.go -f /notdir/
2020/11/04 10:35:26 stat /notdir: no such file or directory
exit status 1
```