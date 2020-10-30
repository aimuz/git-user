package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	_ "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

const useExample = "git user use example"

func (u Users) UseUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use [title]",
		Aliases: []string{"u"},
		Example: "git user use example",
		Short:   "Switch the current repo git user",
		Long:    "Switch the current repo git user",
		Run: func(cmd *cobra.Command, args []string) {
			writer := io.Writer(os.Stdout)
			_cmd := exec.Command("git", "rev-parse", "--show-toplevel")
			b, err := _cmd.CombinedOutput()
			if err != nil {
				_, _ = writer.Write(b)
				return
			}
			gitRepoPath := string(bytes.Trim(b, "\n"))
			if len(args) == 0 {
				_, _ = fmt.Fprintln(writer, "args is empty")
				return
			}
			user, ok := u[args[0]]
			if !ok {
				_, _ = fmt.Fprintf(writer, "%s is not found\n", args[0])
				return
			}
			setConfig := func() error {
				subCmdArgs := map[string][]string{
					"user.name":  {"config", "user.name", user.Name},
					"user.email": {"config", "user.email", user.Email},
				}
				if len(user.IdentityFile) > 0 {
					subCmdArgs["core.sshCommand"] = []string{"config", "core.sshCommand", fmt.Sprintf(`ssh -i %s`, user.IdentityFile)}
				}
				var cmd *exec.Cmd
				for _, strings := range subCmdArgs {
					cmd = exec.Command("git", strings...)
					cmd.Env = os.Environ()
					b, err = cmd.CombinedOutput()
					if err != nil {
						fmt.Println(b, err)
						return err
					}
				}
				return nil
			}
			clean()
			err = setConfig()
			if err != nil {
				_, _ = fmt.Fprintln(writer, err)
				return
			}
			_, _ = fmt.Fprintf(writer,
				"%s setup successfully\n",
				path.Base(gitRepoPath),
			)
		},
	}
	return cmd
}
