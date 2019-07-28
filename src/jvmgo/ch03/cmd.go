package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	XjreOption string
	class string
	args []string
}
/**
	java 命令行格式如下：
		java [-Options] class [args]
		java [-Options] -jar jarFile [args]

	首先设置flag.Usage的值，把printUsage函数赋值给它，然后调用flag包提供的各种Var函数设置需要解析的选项
	接着调用Parse函数解析选项，如果Parse函数解析失败，他就调用printUsage函数把命令的用法打印到控制台
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag,"help",false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) >0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
func printUsage() {
	fmt.Printf("Usage: .%s[-Options] class [args...]\n",os.Args)
}

