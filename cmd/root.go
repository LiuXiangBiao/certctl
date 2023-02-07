/*
Copyright © 2023 lxb
*/
package cmd

import (
	"certificate/logic/automation"
	"os"

	"github.com/spf13/cobra"
)

var automationExample = `
## certcrl use example

./certctl

ps: certctl  -f /root/domains  -a ...id  -s ...secret 
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "certctl",
	Short:   "A tool that detects the generation of certificates",
	Long:    `This tool detects the certificate expiration time, uses the local certbot tool to generate the corresponding certificate, uploads it to Alibaba Cloud SSL App Service, and then notifies the deployment`,
	Version: "0.1.0",
	Example: automationExample,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		automation.CronCertificate()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.certificate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
