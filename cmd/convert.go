/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
  "log"
  "os"

	"github.com/spf13/cobra"
  "github.com/mooijtech/go-pst/v6/pkg"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert PST to PDF",
	Long: `Convert a single PST or multiple into a single PDF or multiple
  PDFs`,
	Run: func(cmd *cobra.Command, args []string) {
    //get input and validate values are supplied
    input, err := cmd.Flags().GetString("input")
    if err != nil {
      log.Fatal(err)
    }

    if input == "" {
      log.Fatal("--input flag is required")
    }

    output, err := cmd.Flags().GetString("output")
    if err != nil {
      log.Fatal(err)
    }

    if output == "" {
      log.Fatal("--output flag is required")
    }

    //check to see if input and output exists
    inputExists, err := exists(input) 
    if err != nil {
      log.Fatal(err)
    }

    if(!inputExists) {
      log.Fatalf("Input %s not found.", input)
    }

    outputExists, err := exists(output)
    if err != nil {
      log.Fatal(err)
    }

    if(!outputExists) {
      log.Fatalf("Output %s not found", output)
    }

    log.Println("Input and output validated. Ready to process.")
    processPST(input, output)
    log.Println("Processing complete.")
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

func processPST(input string, output string) {
  log.Println("Analyzing PST...")
  reader, err := os.Open(input)
  if err != nil {
    log.Fatal(err)
  }

  pstFile, err := pst.New(reader)
  if err != nil {
    log.Fatal(err)
  }

  defer func() {
    pstFile.Cleanup()

    if errClosing := reader.Close(); errClosing != nil {
      log.Fatal(err)
    }
  }()
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

  convertCmd.PersistentFlags().String("input", "", "--input Path to pst or folder")
  convertCmd.PersistentFlags().String("output", "", "--output Path to output folder")
}
