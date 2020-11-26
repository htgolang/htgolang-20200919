To Implement
===

## Protocol
- Server/Client
WriteHeadLen()
ReadHeadLen()
WriteHeadBody()
ReadHeadBody()


## Transport data
- Server:

cmd.Cmd == "ls"  -->  ListFiles()/SendStrContent()
cmd.Cmd == "put" -->  ReceiveFile() 
cmd.Cmd == "get" -->  SendFIle()
cmd.Cmd == "rm"  -->  RemoveFile()

- Client: 

cmd.Cmd == "ls"  -->  ReceiveStrContent()
cmd.Cmd == "put" -->  SendFile() 
cmd.Cmd == "get" -->  ReceiveFIle()
cmd.Cmd == "rm"  -->  ReceiveStrContent()

## data transfer conn

- Server

```golang
// listfile and send content
if cmd.Cmd == "ls" {
    var files []string
    var path = cmd.FilePath
    if cmd.Path != "/" {
        path =  filepath.Join(cmd.FilePath, cmd.FileName)
    } 
    flies = ListFiles(c net.Conn, path)
    SendStrContent(files, len(files))    
}

if cmd.Cmd == "put" {

}
```