package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var port int

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "cmdb web",
	Long:  "cmdb program web",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("web:", verbose, port)
		return nil
	},
}

func init() {
	webCmd.Flags().IntVarP(&port, "port", "P", 8080, "web port")
	rootCmd.AddCommand(webCmd)
}
