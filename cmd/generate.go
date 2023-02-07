/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"certificate/logic/individually"
	"github.com/spf13/cobra"
)

var generateExample = `
# generate certificate example

certctl generate [aksk配置文件] [域名]

`

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a certificate",
	Long:    `此命令使用certbot工具生成单个域名证书`,
	Example: generateExample,
	Run: func(command *cobra.Command, args []string) {
		individually.GenerateCertificate(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
