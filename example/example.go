package main

import (
	"fmt"
	"os"
	"os/exec"

	"bytecode"
)

func main() {
	fmt.Println(bytecode.DescriptorToStackSize("(II[Ljava/lang/String;Z)V"))

	visitor := bytecode.NewClass(bytecode.Java5, "HelloWorld", bytecode.AccPublic|bytecode.AccSuper)
	init := visitor.NewMethod(bytecode.AccPublic, "<init>", "()V")
	init.AddLoadInsn(bytecode.Aload0)
	init.AddMethodInsn(bytecode.Invokespecial, "java/lang/Object",
		"<init>", "()V", false)
	init.AddInsn(bytecode.Return)
	init.End()

	method := visitor.NewMethod(bytecode.AccPublic|bytecode.AccStatic|bytecode.AccVarargs,
		"main", "([Ljava/lang/String;)V")
	method.AddFieldInsn(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	method.AddLdc(bytecode.TypeString, "Hello, World!")
	method.AddMethodInsn(bytecode.Invokevirtual, "java/io/PrintStream",
		"println", "(Ljava/lang/String;)V", false)
	method.AddInsn(bytecode.Return)
	method.End()

	RunJavap(visitor.AsBytecode(), "build/HelloWorld.class")
}

func RunJavap(bytecode []byte, path string) {
	fmt.Println(bytecode)

	file, _ := os.Create(path)
	_, _ = file.Write(bytecode)
	_ = file.Close()

	out, _ := exec.Command("javap", "-verbose", "-c", "-p", path).Output()
	fmt.Println(string(out))
}
