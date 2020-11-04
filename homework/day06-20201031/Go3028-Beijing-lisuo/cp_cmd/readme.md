Demo
===

## Usage
`go run cp_cmd.go --src /path/to/file --dest /path/to/dir`

## usage demo

```bash
root@[10:24:27]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/cp_cmd]
# go run cp_cmd.go --src /path/file --dest /path/
2020/11/04 10:24:56 stat /path/file: no such file or directory
exit status 1
root@[10:24:56]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/cp_cmd]
# go run cp_cmd.go --src readme.md --dest /path/
2020/11/04 10:25:05 stat /path: no such file or directory
exit status 1
root@[10:25:05]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/cp_cmd]
# go run cp_cmd.go --src readme.md --dest /opt/
srcFile:  &{0xc00004e1e0}
destFile type: *os.File, value: &{0xc00004e2a0}: 
root@[10:25:35]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/cp_cmd]
# go run cp_cmd.go --src readme.md --dest /opt/
srcFile:  &{0xc00004e1e0}
There's also a file named readme.md in /opt
Are you want to overwrite it?(y/n): n
Nothing changed.
```