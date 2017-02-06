package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	XJreOption  string
	Class       string
	Args        []string
}

func ParseCmd() *Cmd {
	cmd := new(Cmd)
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "h", false, "print help message.")
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message.")
	flag.BoolVar(&cmd.VersionFlag, "v", false, "print help message.")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print help message.")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XJreOption, "xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
