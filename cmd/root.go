package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "taskify",
	Short: "taskify is a CLI task manager",
}
