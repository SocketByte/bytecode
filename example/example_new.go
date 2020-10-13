package main

import (
	"bytecode"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	visitor := bytecode.NewClass(bytecode.Java5,
		"HelloWorld", "java/lang/Object",
		bytecode.AccPublic|bytecode.AccSuper)

	init := visitor.NewMethod(bytecode.AccPublic, "<init>", "()V")
	init.AddInsn(bytecode.Aload0)
	init.AddMethodInsn(bytecode.Invokespecial,
		"java/lang/Object", "<init>", "()V")
	init.AddInsn(bytecode.Return)

	method := visitor.NewMethod(bytecode.AccPublic|bytecode.AccStatic|bytecode.AccVarargs,
		"main", "([Ljava/lang/String;)V")
	method.AddFieldInsn(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	method.AddLdcInsn(bytecode.TypeString, "Hello, World!")
	method.AddMethodInsn(bytecode.Invokevirtual,
		"java/io/PrintStream", "println", "(Ljava/lang/String;)V")
	method.AddInsn(bytecode.Return)

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
