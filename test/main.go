package main

import (
    "bytes"
    "fmt"
    "log"
    "os/exec"
)

func main() {
    cmd := exec.Command("./cmd/hello.exe", "hah", "zhangsan")
    var stdout, stderr, stdin bytes.Buffer
    stdin.WriteString(`here is stdin
    second line
    third line`)


    cmd.Stdin = &stdin   // 标准输入

    cmd.Stdout = &stdout  // 标准输出
    cmd.Stderr = &stderr  // 标准错误
    err := cmd.Run()
    fmt.Printf("out:\n%s\nerr:\n%s\n", stdout.String(), stdout.String())
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
}