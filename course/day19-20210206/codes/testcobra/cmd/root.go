package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "cmdb",
	Short: "cmdb",
	Long:  "cmdb programe",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("test:", verbose)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}

func init() {
	// 短参数名和长参数名
	// --verbose -V
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "verbose info")
}
