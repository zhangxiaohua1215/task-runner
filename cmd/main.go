package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// 获取命令行参数，第一个参数是程序名，所以从第二个参数开始获取
	args := os.Args[1:]
	// 打印 hello 和参数
	fmt.Println("arg:", args)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	time.Sleep(5 * time.Second)

	fmt.Println("wordcount:", wordCount)

}
