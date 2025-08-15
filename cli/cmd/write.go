//go:build timeserver

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"NovusTimeServer/axi"
	"fmt"

	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "axi write to fpga",
	Long:  `this command is for writing data to the fpga over the uart / axi bus.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("write called")

		op := "write"

		axi.Operate(&op, &args[0], &args[1], &args[2])

	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
