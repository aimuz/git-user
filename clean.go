package main

import "github.com/spf13/cobra"

const clearExample = "git user create --title example --user example --email example@example.com"

func (u Users) CleanUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clear",
		Example: clearExample,
		Short:   "Clear current repo username and email configuration",
		Long: `Clear current repo username and email configuration,
This clearing will cause the commit information to use the configuration of "$HOME/.gitconfig"`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
