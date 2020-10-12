package main

import (
	"fmt"
	"os"
	"os/exec"

	"bytecode"
)

func main() {
	visitor := bytecode.NewClass(bytecode.Java5, "HelloWorld", bytecode.AccPublic|bytecode.AccSuper)
	method := visitor.NewMethod(bytecode.AccPublic, "main", "([Ljava/lang/String;)V")
	method.MaxStackLocals(1, 1)
	method.PushText("Hello, World!")
	method.AddInsn(bytecode.Astore1)
	method.AddInsn(bytecode.Bipush, 120)
	method.AddInsn(bytecode.Istore2)
	method.End()

	method = visitor.NewMethod(bytecode.AccStatic|bytecode.AccPublic, "add", "(II)I")
	method.MaxStackLocals(1, 2)
	method.AddInsn(bytecode.Iload1)
	method.AddInsn(bytecode.Iload2)
	method.AddInsn(bytecode.Iadd)
	method.AddInsn(bytecode.Ireturn)
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
