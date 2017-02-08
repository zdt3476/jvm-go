package main

import (
	"fmt"
	"strings"

	"github.com/zdt3476/jvm-go/classfile"
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
	cn := strings.Replace(cmd.Class, ".", "/", -1)

	cf := loadClass(cn, cp)
	fmt.Println(cmd.Class)

	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("Constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("Access flag: %X\n", cf.AccessFlag())
	fmt.Printf("This class: %X\n", cf.ClassName())
	fmt.Printf("Super class: %X\n", cf.SuperClassName())
	fmt.Printf("Interface names: %X\n", cf.InterfaceNames())
	fmt.Printf("Fields count: %X\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("Field name:%v", f.Name())
	}
	fmt.Printf("Methods count: %X\n", len(cf.Methods()))
	for _, f := range cf.Methods() {
		fmt.Printf("Field name:%v", f.Name())
	}
}
