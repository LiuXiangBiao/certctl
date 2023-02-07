/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"certificate/logic/individually"
	"github.com/spf13/cobra"
)

var checkExample = `
# check certificate command

certvtl check 域名...

`

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "Detection certificate",
	Long:    `此命令可以检查域名证书详细到期时间`,
	Example: checkExample,
	Run: func(cmd *cobra.Command, args []string) {
		individually.CheckCertificate(args)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
