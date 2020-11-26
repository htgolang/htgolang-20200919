To Implement
===

## Protocol
- Server/Client

```golang
WriteHeadLen()
ReadHeadLen()
WriteHeadBody()
ReadHeadBody()
```


## Transport data
- Server/Client:

```golang
// protocol detail will handled in handle func
WriteHeadLen()
ReadHeadLen()
WriteHeadBody()
ReadHeadBody()

cmd.Cmd == "ls"  -->  HandleLS()
cmd.Cmd == "put" -->  HandlePUT()
cmd.Cmd == "get" -->  HandleGET()
cmd.Cmd == "rm"  -->  HandleRM()
```