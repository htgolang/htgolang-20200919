# 作业2

```
PS D:\go_project\golang_practice\day06\go3044mahongbin\2tree> .\tree.exe -h

Usage: ls [OPTION]... [FILE|DIR]...
Options:
  - string
        File/Dir path (default ".")
  -h    Help Info
PS D:\go_project\golang_practice\day06\go3044mahongbin\2tree> .\tree.exe
Args is empty, listing PWD...
↑__ ./
        ↑__ README*
        ↑__ dir0/
                ↑__ 111*
                ↑__ dir1/
                        ↑__ 222*
                        ↑__ dir2/
                                ↑__ file3.cgo*
                        ↑__ file2.ggo*
                ↑__ file1.go*
                ↑__ zzz/
        ↑__ tree.go*
        ↑__ treeRegx.go*
PS D:\go_project\golang_practice\day06\go3044mahongbin\2tree> .\tree.exe dir0 tree.go
↑__ dir0/
        ↑__ 111*
        ↑__ dir1/
                ↑__ 222*
                ↑__ dir2/
                        ↑__ file3.cgo*
                ↑__ file2.ggo*
        ↑__ file1.go*
        ↑__ zzz/
↑__ tree.go*
```
# 作业3 正则版本

```
PS D:\go_project\golang_practice\day06\go3044mahongbin\2+3+4tree> go run .\treeRegx.go dir0
↑__ dir0\

        ↑__ dir1\

                ↑__ dir2\
                        ↑__ file3.cgo*

        ↑__ file1.go*
        ↑__ zzz\
```

# 作业4
```
PS D:\go_project\golang_practice\day06\go3044mahongbin\2+3+4tree> go run .\treeRegxCountLines.go .\dir0\
↑__ .\dir0\\

        ↑__ dir1\

                ↑__ dir2\
                        ↑__ file3.cgo* Lines: 5

        ↑__ file1.go* Lines: 1
        ↑__ zzz\
```
# 作业5

userList.csv
```
48	|	adm	|	4297F44B13955235245B2497399D7A93	|	15166668888	|	北京市	|	2020-11-05	|	false	|	true
49	|	mmm	|	2C216B1BA5E33A27EB6D3DF7DE7F8C36	|	13155556666	|	上海市	|	2020-11-05	|	false	|	false
```