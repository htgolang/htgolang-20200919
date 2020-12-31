package models

type Result struct {
	Rank      int
	IP        string
	IPNum     int
	Status    int
	StatusNum int
	Method    string
	MethodNum int
}

var (
	UploadDir  = "/tmp/upload/"
	UploadFile string
	RankLen    = 10
	IpList     []string
	IpRankList = make(map[int]string)
	MethodList []string
	URLList    []string
	StatusList []string
)

func NewResult(rank, ipnum, status, statusnum, methodnum int, ip, method string) Result {
	return Result{
		Rank:      rank,
		IP:        ip,
		IPNum:     ipnum,
		Status:    status,
		StatusNum: statusnum,
		Method:    method,
		MethodNum: methodnum,
	}
}
