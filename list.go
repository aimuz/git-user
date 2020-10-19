package main

import "github.com/spf13/cobra"

func (u Users) ListUserCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "list all git users",
		Long:    "list all git users",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
