// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the current time",
	Long: `You can use the time show command to view the currennt time. For example:
$./ly show
2020-07-03 15:01:059.`,
	Run: ShowTime,
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//ShowTime
func ShowTime(cmd *cobra.Command, args []string) {
	times := time.Now()
	fmt.Println("当前时间：", times)
}
