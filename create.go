package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const createExample = "git user create --title example --user example --email example@example.com"

func (u Users) CreateUserCommand() *cobra.Command {
	var user = User{}
	var title string
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"a"},
		Example: createExample,
		Short:   "create a new git user",
		Long: `create a new git user, 
The mailbox will be used as the unique identifier`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(user.Name) == 0 || len(user.Email) == 0 {
				return errors.New("username or email cannot be empty")
			}
			if len(title) == 0 {
				title = user.Name
			}
			if len(user.IdentityFile) > 0 {
				_, err := os.Stat(user.IdentityFile)
				if err != nil {
					return err
				}
			}
			u[title] = user
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
	flags.StringVar(&user.Name, "user", user.Name, "git user name")
	flags.StringVar(&user.Email, "email", user.Email, "git user email")
	flags.StringVarP(&user.IdentityFile, "identity_file", "i", user.IdentityFile,
		"The certificate corresponding to the user. If it is blank, the default value will be used")
	return cmd
}
