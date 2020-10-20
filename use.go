package main

import (
	"fmt"
	_ "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func (u Users) UseUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use",
		Aliases: []string{"u"},
		Short:   "Switch the current repo git user",
		Long:    "Switch the current repo git user",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args)
			return nil
		},
	}
	return cmd
}
