package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "golf",
	Short: "Golf score tracking CLI",
}
