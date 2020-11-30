package socket

const (
	proto = "tcp"
	addr  = ":8081"
)

// ResponseHead  represents operation and status
type ResponseHead struct {
	Operation string
	Status    int
}
