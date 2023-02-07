/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"certificate/logic/individually"
	"fmt"
	"github.com/spf13/cobra"
)

var uploadExample = `
# upload certificate example

certctl upload [accessKeyId] [accessKeySecret] [certname] [cert] [key]

accessKeyId： 阿里云账号的ak
accessKeySecret： 阿里云账号的sk
certname： 设置此证书名字
cert：证书pem
key：证书私钥

`

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "Upload the certificate",
	Long:    `此命令把已生成的证书上传到阿里云ssl应用服务上`,
	Example: uploadExample,
	Run: func(cmd *cobra.Command, args []string) {
		err := individually.UploadCertificate(args[0], args[1], args[2], args[3], args[4])
		if err != nil {
			fmt.Printf("upload certificate faield", err)
		} else {
			fmt.Println("success")
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
