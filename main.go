package main

import (
	"fmt"
	"os"
)

func readLinesParallel(fileName string, ch chan string) {
	content, err := os.ReadFile(fileName)
	if err == nil {
		ch <- string(content)
	}
}

func printData(data string, indexStr string) {
	fmt.Printf("-------Start Printing from %s goroutine-------", indexStr)
	fmt.Println(data)
	fmt.Println("---------END-----------")
	fmt.Println()
}

func main() {
	chMap := make(map[string]chan string)
	fileNames := [3]string{"a.txt", "b.txt", "c.txt"}
	for _, f := range fileNames {
		chMap[f] = make(chan string)
		go readLinesParallel(f, chMap[f])
	}
	i := 3
	for i > 0 {
		select {
		case data := <-chMap["a.txt"]:
			printData(data, "1st")
			i = i - 1
		case data := <-chMap["b.txt"]:
			printData(data, "2nd")
			i = i - 1
		case data := <-chMap["c.txt"]:
			printData(data, "3rd")
			i = i - 1
		}
	}

	fmt.Println("Ending printing")
}
