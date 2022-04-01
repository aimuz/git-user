package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

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
					"user.name":  {user.Name},
					"user.email": {user.Email},
				}
				if len(user.IdentityFile) > 0 {
					subCmdArgs["core.sshCommand"] = []string{fmt.Sprintf(`ssh -i %s`, user.IdentityFile)}
				}

				if len(user.GPGKey) > 0 {
					subCmdArgs["user.signingkey"] = []string{user.GPGKey}
				}

				var gitSubCmd *exec.Cmd
				for subCmd, cmdArgs := range subCmdArgs {
					gitSubCmd = exec.Command("git", append([]string{"config", "--local", subCmd}, cmdArgs...)...)
					gitSubCmd.Env = os.Environ()
					b, err = gitSubCmd.CombinedOutput()
					if err != nil {
						fmt.Println(string(b), err)
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
