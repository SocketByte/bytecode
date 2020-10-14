# Bytecode
Very simple bytecode generation library for Go

## Features
- Creating Java classes, methods and global fields
- Jump labels/instructions
- Automatic stack depth calculation
- Automatic locals calculation
- Efficient generation
- Low-level access
- Simple and self-documented codebase
- Constant pool caching

## Missing features (yet)
- Annotation support
- `LineNumberTable` support
- Exception support (`throws`)
- Interface support
- Inner classes
- Stack frames

## Basic example
```go
package main

import (
    "fmt"
    "github.com/SocketByte/bytecode"
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

    fmt.Println(visitor.AsBytecode())
}

```
```
public class HelloWorld
  minor version: 0
  major version: 49
  flags: (0x0021) ACC_PUBLIC, ACC_SUPER
  this_class: #24                         // HelloWorld
  super_class: #4                         // java/lang/Object
  interfaces: 0, fields: 2, methods: 2, attributes: 1
Constant pool:
   #1 = Utf8               SourceFile
   #2 = Utf8               HelloWorld.java
   #3 = Utf8               java/lang/Object
   #4 = Class              #3             // java/lang/Object
   #5 = Utf8               <init>
   #6 = Utf8               ()V
   #7 = NameAndType        #5:#6          // "<init>":()V
   #8 = Methodref          #4.#7          // java/lang/Object."<init>":()V
   #9 = Utf8               java/lang/System
  #10 = Class              #9             // java/lang/System
  #11 = Utf8               out
  #12 = Utf8               Ljava/io/PrintStream;
  #13 = NameAndType        #11:#12        // out:Ljava/io/PrintStream;
  #14 = Fieldref           #10.#13        // java/lang/System.out:Ljava/io/PrintStream;
  #15 = Utf8               Hello, World!
  #16 = String             #15            // Hello, World!
  #17 = Utf8               java/io/PrintStream
  #18 = Class              #17            // java/io/PrintStream
  #19 = Utf8               println
  #20 = Utf8               (Ljava/lang/String;)V
  #21 = NameAndType        #19:#20        // println:(Ljava/lang/String;)V
  #22 = Methodref          #18.#21        // java/io/PrintStream.println:(Ljava/lang/String;)V
  #23 = Utf8               HelloWorld
  #24 = Class              #23            // HelloWorld
  #25 = Utf8               Code
  #26 = Utf8               main
  #27 = Utf8               ([Ljava/lang/String;)V
  #28 = Utf8               globalValue
  #29 = Utf8               I
  #30 = Integer            43
  #31 = Utf8               ConstantValue
  #32 = Utf8               globalString
  #33 = Utf8               Ljava/lang/String;
{
  public final int globalValue;
    descriptor: I
    flags: (0x0011) ACC_PUBLIC, ACC_FINAL
    ConstantValue: int 43

  public java.lang.String globalString;
    descriptor: Ljava/lang/String;
    flags: (0x0001) ACC_PUBLIC

  public HelloWorld();
    descriptor: ()V
    flags: (0x0001) ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #8                  // Method java/lang/Object."<init>":()V
         4: return

  public static void main(java.lang.String...);
    descriptor: ([Ljava/lang/String;)V
    flags: (0x0089) ACC_PUBLIC, ACC_STATIC, ACC_VARARGS
    Code:
      stack=2, locals=2, args_size=1
         0: bipush        -3
         2: istore_1
         3: iload_1
         4: iconst_2
         5: if_icmpge     16
         8: getstatic     #14                 // Field java/lang/System.out:Ljava/io/PrintStream;
        11: ldc           #16                 // String Hello, World!
        13: invokevirtual #22                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
        16: return
}
SourceFile: "HelloWorld.java"
```