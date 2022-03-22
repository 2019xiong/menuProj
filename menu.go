package main

import "fmt"

func main() {
	var str string
	for {
		fmt.Print("请输入命令：")
		fmt.Scan(&str)
		if str == "quit" {
			break
		}
		switch str {
		case "help":
			fmt.Println("this is a menu project")
			fmt.Println("Usage:")
			fmt.Println("\t<command>")
			fmt.Println("The commands are: ....")
			fmt.Println("\tlist\tlist something")
			fmt.Println("\tquit\t输入quit退出")
		case "list":
			fmt.Println("some list information")
		default:
			fmt.Println("unknown command")
			fmt.Println("Run: 'help' for usage.")
		}
	}
}
