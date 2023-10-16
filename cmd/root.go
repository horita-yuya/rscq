/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/horita-yuya/rscq/cmd/react"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "rscq",
	Short: "Simple parser for React Server Components",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			r := os.Stdin
			b, _ := io.ReadAll(r)
			lines := react.ParseRSC(string(b))
			react.DisplayRSC(lines)
		}
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.SetOut(color.Output)
}
