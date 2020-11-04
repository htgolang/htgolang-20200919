Demo
===

## Usage 

```bash
root@[10:43:13]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/code_lines]
# go run code_lines.go -f /opt/
file: 2.go                        dir: false
file: 3.cgo                       dir: false
all the files in dir /opt end with .go&.cgo has 7 lines
root@[10:43:18]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/code_lines]
# go run code_lines.go -f ../user_manager_proj/
file: main.go                     dir: false
all the files in dir /data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/user_manager_proj end with .go&.cgo has 40 lines
root@[10:43:34]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/code_lines]
# go run code_lines.go -f ../user_manager_proj/main.go 
file: main.go                     dir: false
all the files at /data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/user_manager_proj end with .go&.cgo has 40 lines
root@[10:43:37]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/code_lines]
# go run code_lines.go -f ../user_manager_proj/cmd/funcs/
file: adduser.go                  dir: false
file: cmd_func_map.go             dir: false
file: deluser.go                  dir: false
file: login.go                    dir: false
file: misc.go                     dir: false
file: moduser.go                  dir: false
file: queryuser.go                dir: false
file: show_user_list.go           dir: false
file: user_op.go                  dir: false
all the files in dir /data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/user_manager_proj/cmd/funcs end with .go&.cgo has 778 lines
root@[10:43:47]suosuoli:[/data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/code_lines]
# go run code_lines.go -f ../user_manager_proj/cmd/db/
file: define.go                   dir: false
file: readuser.go                 dir: false
file: saveuser.go                 dir: false
all the files in dir /data/htgolang-20200919/homework/day06-20201031/Go3028-Beijing-lisuo/user_manager_proj/cmd/db end with .go&.cgo has 113 lines
```