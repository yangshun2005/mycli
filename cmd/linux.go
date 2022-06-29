/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// linuxCmd represents the linux command
var linuxCmd = &cobra.Command{
	Use:   "linux",
	Short: "L2 command linux",
	Long:  `Linux Command for mycli as L2`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mycli linux <Linux CMD>")
	},
}

func init() {
	rootCmd.AddCommand(linuxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linuxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linuxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
