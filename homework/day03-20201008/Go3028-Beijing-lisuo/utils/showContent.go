package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func ShowContent(id int64, name, address, phone string) {
	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, string(id)+"\t"+name+"\t"+phone+"\t"+address)
	w.Flush()
}
