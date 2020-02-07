/*
Copyright © 2020 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/victor23d/vault-toolkit/api"
	"go.uber.org/zap"
)

var (
logger, _      = zap.NewProduction()
log            = logger.Sugar()
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Print All Secrets
		secrets := api.GetAllSecret()
		fmt.Println("================================================================================")
		fmt.Println("================================================================================")

		secretJson, err := json.MarshalIndent(secrets, "", "    ")
		if err != nil {
			log.Error(err)
		}
		fmt.Println(string(secretJson))

	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
