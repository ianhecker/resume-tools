package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	err := rootCmd.Execute()
	checkErr(err)
}

func init() {
	rootCmd.AddCommand(skillCmd)
}

func readFile(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("could not read file: %w", err)
		checkErr(err)
	}
	return bytes, nil
}

func writeFile(filename string, bytes []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(bytes))
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
