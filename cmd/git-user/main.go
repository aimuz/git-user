package main

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var dataPath = os.Getenv("HOME") + "/.config/git-user/user.yaml"

var users = Users{}

func main() {
	rootCmd := &cobra.Command{
		Use:     "user [command] [flags]",
		Short:   "git multi user management",
		Example: strings.Join([]string{createExample, listExample, useExample}, "\n "),
		RunE:    run,
		Version: "v0.1.0",
	}
	setUsageTemplate(rootCmd)
	globalFlags := rootCmd.PersistentFlags()
	globalFlags.StringVar(&dataPath, "data", dataPath, ``)
	if err := initCommand(rootCmd); err != nil {
		panic(err)
	}
	_ = rootCmd.Execute()
}

func initCommand(cmd *cobra.Command) error {
	dir := path.Dir(dataPath)
	_, err := os.Open(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
	}
	if err != nil {
		return err
	}
	file, err := os.OpenFile(dataPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	err = yaml.NewDecoder(file).Decode(&users)
	if err != nil && err != io.EOF {
		err = nil
	}
	cmd.AddCommand(
		users.CreateUserCommand(),
		users.ListUserCommand(),
		users.UseUserCommand(),
		users.CleanUserCommand(),
	)
	return nil
}

func run(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func setUsageTemplate(cmd *cobra.Command) {
	cmd.SetUsageTemplate(`Usage:{{if .Runnable}}
 git {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
 git {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
 {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
 {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
 git user {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
 git {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)
}
