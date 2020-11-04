Demo
===

## Usage

```bash
root@[10:27:35]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/get_file_dir]
# go run get_file_dir.go -f /opt/
file: 2.go                        dir: false
file: 3.cgo                       dir: false
file: Platform-projs-notes        dir: true
file: check.py                    dir: false
file: containerd                  dir: true
file: dex.xml                     dir: false
file: dex.xmln                    dir: false
file: go                          dir: false
file: mongodb-3.2.10              dir: true
file: mongodb_conf                dir: true
file: online-3.2.10.zip           dir: false
file: readme.md                   dir: false
file: src                         dir: false
root@[10:27:41]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/get_file_dir]
# go run get_file_dir.go -f /opt/dex.xml
file: 2.go                        dir: false
file: 3.cgo                       dir: false
file: Platform-projs-notes        dir: true
file: check.py                    dir: false
file: containerd                  dir: true
file: dex.xml                     dir: false
file: dex.xmln                    dir: false
file: go                          dir: false
file: mongodb-3.2.10              dir: true
file: mongodb_conf                dir: true
file: online-3.2.10.zip           dir: false
file: readme.md                   dir: false
file: src                         dir: false
root@[10:27:50]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/get_file_dir]
# go run get_file_dir.go -f /noSuchDir
2020/11/04 10:28:04 stat /noSuchDir: no such file or directory
exit status 1
root@[10:28:04]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/get_file_dir]
# go run get_file_dir.go -f /opt/noSuchfile
2020/11/04 10:28:12 stat /opt/noSuchfile: no such file or directory
exit status 1
root@[10:32:18]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/get_file_dir]
# go run get_file_dir.go 
Usage: go run get_file_dir.go -f /path/to/fileOrDir
2020/11/04 10:32:26 stat /data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/get_file_dir/file-or-dir: no such file or directory
exit status 1
```