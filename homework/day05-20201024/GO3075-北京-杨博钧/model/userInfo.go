package model

import (
	"github.com/olekukonko/tablewriter"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id int
	Name string
	Addr string
	Tel string
	Birthday time.Time
	Password string
}
//以表格的形式打印单条用户信息

func (this User) String() string {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Id", "Name", "Addr", "Tel", "Birthday", "Password"})
	table.Append([]string{strconv.Itoa(this.Id), this.Name, this.Addr, this.Tel,
		this.Birthday.Format("2006-01-02"), this.Password })
	table.Render()
	return tableString.String()
}
