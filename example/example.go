package main

import (
	"fmt"
	"os"
	"os/exec"

	"bytecode"
)

func main() {
	fmt.Println(fmt.Sprintf("%x", 29343))

	visitor := bytecode.NewClass(bytecode.Java5, "HelloWorld", bytecode.AccPublic|bytecode.AccSuper)
	init := visitor.NewMethod(bytecode.AccPublic, "<init>", "()V")
	init.AddInsn(bytecode.Aload0)
	init.AddMethodInsn(bytecode.Invokespecial, "java/lang/Object", "<init>", "()V")
	init.End()

	method := visitor.NewMethod(bytecode.AccPublic|bytecode.AccStatic, "main", "([Ljava/lang/String;)V")
	method.MaxStackLocals(1, 1)
	method.AddLdc(bytecode.TypeString, "siema")
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
