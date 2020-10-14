package main

import (
    "bytecode"
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // Create new Java class
    visitor := bytecode.NewClass(bytecode.Java5,
        "HelloWorld", "java/lang/Object",
        bytecode.AccPublic|bytecode.AccSuper)

    // Add source file attribute
    visitor.AddSourceFile("HelloWorld.java")

    // Append 2 global fields
    visitor.NewField(bytecode.AccPublic|bytecode.AccFinal,
        "globalValue", "I", 43)
    visitor.NewField(bytecode.AccPublic,
        "globalString", "Ljava/lang/String;", nil)

    // Create new basic constructor
    init := visitor.NewMethod(bytecode.AccPublic, "<init>", "()V")
    init.AddInsn(bytecode.Aload0)
    init.AddMethodInsn(bytecode.Invokespecial,
        "java/lang/Object", "<init>", "()V")
    init.AddInsn(bytecode.Return)

    // Create main method
    main := visitor.NewMethod(bytecode.AccPublic|bytecode.AccStatic|bytecode.AccVarargs,
        "main", "([Ljava/lang/String;)V")
    // Push -3 as byte value
    main.AddInt8Insn(bytecode.Bipush, -3)
    // Store at 1
    main.AddVarInsn(bytecode.Istore1)
    // Load 1
    main.AddInsn(bytecode.Iload1)
    // Push 2 as int value
    main.AddInsn(bytecode.Iconst2)
    // Create a new jump label
    label := main.NewLabel()
    // Add jump instruction (control-flow) to a newly created label
    main.AddJumpInsn(bytecode.Ificmpge, label)
    // System.out.println("Hello, World!");
    main.AddFieldInsn(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
    main.AddLdcInsn(bytecode.Ldc, "Hello, World!")
    main.AddMethodInsn(bytecode.Invokevirtual,
        "java/io/PrintStream", "println", "(Ljava/lang/String;)V")
    // Add return to the method
    main.AddInsn(bytecode.Return)
    // Set label to point to "return" instruction (jump to return)
    main.AddLabel(label)

    RunJavap(visitor.AsBytecode(), "build/HelloWorld.class")
}

func RunJavap(bytecode []byte, path string) {
    file, _ := os.Create(path)
    _, _ = file.Write(bytecode)
    _ = file.Close()

    out, _ := exec.Command("javap", "-verbose", "-c", "-p", path).Output()
    fmt.Println(string(out))
}
