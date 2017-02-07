package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zdt3476/jvm-go/classpath"
	jcmd "github.com/zdt3476/jvm-go/cmd"
)

func main() {
	cmd := jcmd.ParseCmd()

	if cmd.VersionFlag {
		fmt.Println("Version 0.0.1 beta.")
	} else if cmd.HelpFlag || len(cmd.Class) <= 0 {
		jcmd.PrintUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *jcmd.Cmd) {
	cp := classpath.Parse(cmd.XJreOption, cmd.CpOption)
	fmt.Printf("Jvm: classpath[%v] class[%s] args[%v]\n", cp, cmd.Class, cmd.Args)

	cn := strings.Replace(cmd.Class, ".", string(os.PathSeparator), -1)
	data, _, err := cp.ReadClass(cn)
	if err != nil {
		log.Fatalf("Can not found or load main class:%v\n", cmd.Class)
	}

	fmt.Printf("This class is: %v\n", string(data))
}
