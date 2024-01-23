/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "log"
  "os"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert PST to PDF",
	Long: `Convert a single PST or multiple into a single PDF or multiple
  PDFs`,
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
      log.Fatal("Must provide a folder or file to convert.")
    }

    

    reader, err := os.Open(args[0])
    if err != nil {
      log.Fatalf()
    }
  },
}

func exists(path string) (bool, error){
  _, err := os.Stat(path)
  if err == nil {
    return true, nil
  }

  if os.IsNotExist(err) {
    return false, nil
  }

  return false, err
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
