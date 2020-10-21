package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

func (u Users) ListUserCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Example: "git user list",
		Short:   "list all git users",
		Long:    "list all git users",
		RunE: func(cmd *cobra.Command, args []string) error {
			w := &tabwriter.Writer{}
			w.Init(os.Stdout, 8, 8, 0, '\t', 0)
			defer w.Flush()
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t\n", "TITLE", "USER", "EMAIL")
			for s, user := range u {
				_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t\n", s, user.Name, user.Email)
			}
			return nil
		},
	}
}
