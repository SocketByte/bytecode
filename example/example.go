package main

import (
    "fmt"
    "os"
    "os/exec"

    "bytecode"
)

func main() {
	visitor := bytecode.NewClass(bytecode.Java5, "HelloWorld", bytecode.AccPublic | bytecode.AccSuper)
	method := visitor.NewMethod(bytecode.AccPublic, "main", "([Ljava/lang/String;)V")
	method.MaxStackLocals(1, 1)
	method.PushText("Hello, World!")
	method.AddInsn(bytecode.Astore1)
	method.End()

	RunJavap(visitor.AsBytecode(), "build/hello.class")
}

func RunJavap(bytecode []byte, path string) {
    fmt.Println(bytecode)

    file, _ := os.Create(path)
    _, _ = file.Write(bytecode)
    _ = file.Close()

    out, _ := exec.Command("javap", "-verbose", "-c", "-p", path).Output()
    fmt.Println(string(out))
}