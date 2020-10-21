package main

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	_ "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
)

func (u Users) UseUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use [title]",
		Aliases: []string{"u"},
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
			if len(args) == 0 {
				_, _ = fmt.Fprintln(writer, "args is empty")
				return
			}
			user, ok := u[args[0]]
			if !ok {
				_, _ = fmt.Fprintf(writer, "%s is not found\n", args[0])
				return
			}
			gitRepoPath := string(bytes.Trim(b, "\n"))
			repo, err := git.PlainOpen(gitRepoPath)
			if err != nil {
				_, _ = fmt.Fprintln(writer, err)
				return
			}
			cfg, err := repo.Config()
			if err != nil {
				_, _ = fmt.Fprintln(writer, err)
				return
			}
			cfg.User.Name = user.Name
			cfg.User.Email = user.Email
			err = repo.SetConfig(cfg)
			if err != nil {
				return
			}
			_, _ = fmt.Fprintln(writer, "User set successfully")
		},
	}
	return cmd
}
