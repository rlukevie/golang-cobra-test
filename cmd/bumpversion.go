/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// bumpversionCmd represents the bumpversion command
var bumpversionCmd = &cobra.Command{
	Use:   "bumpversion",
	Short: "Bump VERSION by 1",
	Long:  `This command reads the VERSION file and increases the saved number by 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		bumpVersion(args)
	},
}

func bumpVersion(args []string) {
	data, err := os.ReadFile("VERSION")
	if err != nil {
		fmt.Println("Could not read from VERSION file.")
		return
	}
	version, err := strconv.Atoi(string(data))
	newVersion := version + 1
	newData := []byte(strconv.Itoa(newVersion))
	err = os.WriteFile("VERSION", newData, 0666)
}

func init() {
	rootCmd.AddCommand(bumpversionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bumpversionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bumpversionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
