package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本显示",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo四则运算生成器v0.1")
	},
}
