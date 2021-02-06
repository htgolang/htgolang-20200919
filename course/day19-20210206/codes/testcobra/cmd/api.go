package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var host string

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "cmdb api",
	Long:  "cmdb program api",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("api:", host, verbose)
		return nil
	},
}

func init() {
	apiCmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "host")
	rootCmd.AddCommand(apiCmd)
}
