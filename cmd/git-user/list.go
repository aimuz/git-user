package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

const listExample = "git user list"

func (u Users) ListUserCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Example: "git user list",
		Short:   "list all git users",
		Long:    "list all git users",
		RunE: func(cmd *cobra.Command, args []string) error {
			w := &tabwriter.Writer{}
			w.Init(os.Stdout, 8, 8, 0, '\t', 0)
			defer w.Flush()
			_, _ = fmt.Fprint(w, "TITLE\tUSER\tEMAIL\tIDENTITY FILE\tGPG KEY\n")
			for s, user := range u {
				title := s
				if user.Default {
					title += "(default)"
				}
				_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", s, user.Name, user.Email, user.IdentityFile)
			}
			return nil
		},
	}
}
