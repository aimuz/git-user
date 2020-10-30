package main

import (
	"flag"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
