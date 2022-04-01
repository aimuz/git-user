package main

import (
	"fmt"
	"os"
	"sort"
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
			keys := make([]string, 0, len(u))
			for key := range u {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			for _, key := range keys {
				title := key
				user := u[key]
				if user.Default {
					title += "(default)"
				}
				_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", title, user.Name, user.Email, user.IdentityFile, user.GPGKey)
			}
			return nil
		},
	}
}
