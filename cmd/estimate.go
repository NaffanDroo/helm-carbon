/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// estimateCmd represents the estimate command
var estimateCmd = &cobra.Command{
	Use:   "estimate",
	Short: "Provide a basic estimate of the carbon foot print",
	Long: `Use the helm chart and values to provide an estimate of the
carbon footprint`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("estimate called")
		// Suppress command usage on error
		cmd.SilenceUsage = true

	},
}

func init() {
	rootCmd.AddCommand(estimateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// estimateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// estimateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
