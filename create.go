package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func (u Users) CreateUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"a"},
		Short:   "create a new git user",
		Long: `create a new git user, 
The mailbox will be used as the unique identifier`,
		RunE: func(cmd *cobra.Command, args []string) error {
			title, _ := cmd.Flags().GetString("title")
			username, _ := cmd.Flags().GetString("user")
			email, _ := cmd.Flags().GetString("email")
			if len(username) == 0 || len(email) == 0 {
				return errors.New("username or email cannot be empty")
			}
			if len(title) == 0 {
				title = username
			}
			u[title] = User{
				Name:  username,
				Email: email,
			}
			err := u.Update()
			if err != nil {
				return err
			}
			fmt.Printf("Successfully created %s user\n", title)
			return nil
		},
	}
	flags := cmd.Flags()
	flags.String("title", "", "if it is empty, username will be used")
	flags.String("user", "", "git user name")
	flags.String("email", "", "git user email")
	return cmd
}
