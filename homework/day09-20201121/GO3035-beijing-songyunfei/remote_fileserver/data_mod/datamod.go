package data_mod

type Cmd struct {
	Cmd string
	Arg string
}

type Data struct {
	Lenght int
	Ack    bool
	Data   []byte
	Error  string
}
