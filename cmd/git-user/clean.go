package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"path"
)

const clearExample = "git user create --title example --user example --email example@example.com"

func (u Users) CleanUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clean",
		Example: clearExample,
		Short:   "Clean current repo username and email configuration",
		Long: `Clean current repo username and email configuration,
This clearing will cause the commit information to use the configuration of "$HOME/.gitconfig"`,
		Run: func(cmd *cobra.Command, args []string) {
			clean()
		},
	}
	return cmd
}

func clean() {
	writer := io.Writer(os.Stdout)
	_cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	b, err := _cmd.CombinedOutput()
	if err != nil {
		_, _ = writer.Write(b)
		return
	}
	gitRepoPath := string(bytes.Trim(b, "\n"))

	setConfig := func() error {
		subCmdArgs := map[string][]string{
			"user.name-temp":       {"config", "user.name", "temp"},
			"user.name":            {"config", "--unset", "user.name"},
			"user.email-temp":      {"config", "user.email", "temp@example.com"},
			"user.email":           {"config", "--unset", "user.email"},
			"core.sshCommand-temp": {"config", "core.sshCommand", "temp"},
			"core.sshCommand":      {"config", "--unset", "core.sshCommand"},
		}
		var cmd *exec.Cmd
		for _, strings := range subCmdArgs {
			cmd = exec.Command("git", strings...)
			_, err = cmd.CombinedOutput()
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = setConfig()
	if err != nil {
		_, _ = fmt.Fprintln(writer, err)
		return
	}
	_, _ = fmt.Fprintf(writer,
		"%s clean successfully\n",
		path.Base(gitRepoPath),
	)
}
