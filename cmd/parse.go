// Package cmd /*
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Format current time",
	Long: `Format current time. For example

time show parse --format "2006/01/02 15:04:05".`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now().Format(format))
	},
}

var format string

func init() {
	showCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringVarP(&format, "format", "f", "", "Help message for toggle")
	_ = parseCmd.MarkFlagRequired("format")
}
