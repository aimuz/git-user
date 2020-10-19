package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
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
			repo, err := git.PlainOpen("./")
			if err != nil {
				return err
			}
			fmt.Println(repo.Config())
			return nil
		},
	}
	return cmd
}
