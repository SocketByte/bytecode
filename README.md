# Bytecode
Very simple bytecode generation library for Go

```go
package main

import (
    "bytecode"
    "fmt"
)

func main() {
    visitor := bytecode.NewClass(bytecode.Java5, "HelloWorld", bytecode.AccPublic|bytecode.AccSuper)
    init := visitor.NewMethod(bytecode.AccPublic, "<init>", "()V")
    init.AddInsn(bytecode.Aload0)
    init.AddMethodInsn(bytecode.Invokespecial, "java/lang/Object", "<init>", "()V", false)
    init.AddInsn(bytecode.Return)
    init.End(1)

    method := visitor.NewMethod(bytecode.AccPublic|bytecode.AccStatic|bytecode.AccVarargs,
        "main", "([Ljava/lang/String;)V")
    method.AddFieldInsn(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
    method.AddLdc(bytecode.TypeString, "Hello, World!")
    method.AddMethodInsn(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(Ljava/lang/String;)V", false)
    method.AddInsn(bytecode.Return)
    method.End(2)

    bytecodeArray := visitor.AsBytecode()
    fmt.Println(bytecodeArray)
}
```