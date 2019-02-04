package cmd

import "github.com/spf13/cobra"

// RegisterableCommand wraps Cobra Commands and allows the inclusion of a register
// function for performing any preparations and then adding it to a root command.
type RegisterableCommand interface {
	registerWith(cmd *cobra.Command)
}
