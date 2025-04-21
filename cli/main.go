package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "golf",
		Short: "Golf score tracking CLI",
	}

	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Says hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Allen")
		},
	}

	rootCmd.AddCommand(helloCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
