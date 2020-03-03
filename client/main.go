package main

import (
	"bufio"
	"fmt"
	"github.com/Planck1858/todo_list_cli/client/API"
	"log"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	clearAll()
	for {
		printMenu()
		s, _, err := reader.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		str := string(s)

		switch str {
		case "a":
			clearAll()

			res := API.TryConnection()

			printSlash()
			fmt.Println(res)
			printSlash()
			fmt.Scan()
		case "b":
			clearAll()

			res := API.GetList()

			printSlash()
			fmt.Println("Tasks:")
			showList(res)
			printSlash()
			fmt.Scan()
		case "c":
			clearAll()
			fmt.Println("Print your task:")
			str, _, err := reader.ReadLine()
			if err != nil {
				log.Fatalln(err)
			}

			API.NewTask(string(str))

			printSlash()
			fmt.Scan()
		case "d":
			clearAll()
			res := API.GetList()
			showList(res)
			fmt.Println("")
			fmt.Println("Choose task [index]:")
			str, _, err := reader.ReadLine()
			if err != nil {
				log.Fatalln(err)
			}

			API.ChangeTask(string(str))

			printSlash()
			fmt.Scan()
		case "e":
			clearAll()
			res := API.GetList()
			showList(res)
			fmt.Println("")
			fmt.Println("Choose task [index]:")
			str, _, err := reader.ReadLine()
			if err != nil {
				log.Fatalln(err)
			}

			API.DeleteTask(string(str))

			printSlash()
			fmt.Scan()
		case "f":
			clearAll()

			API.DeleteList()

			printSlash()
			fmt.Scan()
		case "q":
			os.Exit(3)

		default:
			clearAll()
		}

	}
}

func showList(list API.List) {
	for _, v := range list.Tasks {
		str := fmt.Sprintf("%s) %s - status[%s]", strconv.Itoa(v.Index), v.Task, strconv.FormatBool(v.Completed))
		fmt.Println(str)
	}
}

func printSlash() {
	fmt.Println("////////////////////////////////////////")
}

func printMenu() {
	fmt.Println("")
	fmt.Println("////////////// Menu: //////////////")
	fmt.Println("a) Try Connection")
	fmt.Println("b) Get List")
	fmt.Println("c) New Task")
	fmt.Println("d) Change Task Status")
	fmt.Println("e) Delete Task")
	fmt.Println("f) Clear All List")
	fmt.Println("q) Quit")
	fmt.Println("")
	fmt.Println("Choose command: ")
	fmt.Println("")
}

func clearAll() {
	print("\033[H\033[2J")
}
