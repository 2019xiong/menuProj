package main

import "fmt"

func main() {
	// var showMenu = menu()
	// showMenu()
	testLinkTable()
}

func testLinkTable() {
	linktable := initLinkTable()
	insertTailNode(1, linktable)
	insertTailNode(2, linktable)
	insertTailNode(3, linktable)
	insertTailNode(4, linktable)
	insertTailNode(5, linktable)
	insertTailNode(6, linktable)
	fmt.Println("拥有6个节点的初始链表：")
	showLinkTableData(linktable)
	fmt.Println("在索引2处,插入值为7的节点:")
	append(2, 7, linktable)
	showLinkTableData(linktable)
	fmt.Println("查找索引为1的节点:")
	fmt.Println(findNodeByIndex(1, linktable))
	fmt.Println("查找值为7的节点:")
	fmt.Println(findNodeByValue(7, linktable))
	fmt.Println("删除第6个节点:")
	deleteNode(6, linktable)
	showLinkTableData(linktable)
}

func menu() func() {
	var str string
	fun := func() {
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
	return fun
}
