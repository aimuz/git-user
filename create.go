package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

const createExample = "git user create --title example --user example --email example@example.com"

func (u Users) CreateUserCommand() *cobra.Command {
	var title, username, email string
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"a"},
		Example: createExample,
		Short:   "create a new git user",
		Long: `create a new git user, 
The mailbox will be used as the unique identifier`,
		RunE: func(cmd *cobra.Command, args []string) error {
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
	flags.StringVar(&title, "title", title, "if it is empty, username will be used")
	flags.StringVar(&username, "user", username, "git user name")
	flags.StringVar(&email, "email", email, "git user email")
	return cmd
}
