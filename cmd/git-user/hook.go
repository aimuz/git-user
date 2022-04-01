package main

import "github.com/spf13/cobra"

func (u Users) HookCommand() *cobra.Command {
	return nil
}

func (u Users) Hook() error {
	return nil
}
