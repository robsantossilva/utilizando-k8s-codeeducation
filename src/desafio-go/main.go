package main

import "fmt"

func main() {
	fmt.Println(greeting("Code.education Rocks!"))
}

func greeting(txt string) string {
	return "<b>" + txt + "</b>"
}
