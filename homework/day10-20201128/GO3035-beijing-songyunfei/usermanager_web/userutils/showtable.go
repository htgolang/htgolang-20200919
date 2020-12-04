package userutils

//import (
//"github.com/olekukonko/tablewriter"
//	"io"
//	"os"
//)
//
//func Showintable(header []string, data [][]string){
//	table := tablewriter.NewWriter(os.Stdout)
//	table.SetHeader(header)
//	for _, v := range data {
//		table.Append(v)
//	}
//	table.Render()
//}
//func Gentable(header []string, data [][]string,o io.Writer){
//	table := tablewriter.NewWriter(o)
//	table.SetHeader(header)
//	for _, v := range data {
//		table.Append(v)
//	}
//	table.Render()
//}