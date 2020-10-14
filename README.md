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

    // Append 2 global fields
    visitor.NewField(bytecode.AccPublic, "globalValue", "I")
    visitor.NewField(bytecode.AccPublic, "globalString", "Ljava/lang/String;")

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
    main.AddInsn(bytecode.Istore1)
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
    main.AddLdcInsn(bytecode.TypeString, "Hello, World!")
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
Classfile /C:/Cloud/GoProjects/bytecode/build/HelloWorld.class
public class HelloWorld
  minor version: 0
  major version: 49
  flags: (0x0021) ACC_PUBLIC, ACC_SUPER
  this_class: #22                         // HelloWorld
  super_class: #2                         // java/lang/Object
  interfaces: 0, fields: 2, methods: 2, attributes: 0
Constant pool:
   #1 = Utf8               java/lang/Object
   #2 = Class              #1             // java/lang/Object
   #3 = Utf8               <init>
   #4 = Utf8               ()V
   #5 = NameAndType        #3:#4          // "<init>":()V
   #6 = Methodref          #2.#5          // java/lang/Object."<init>":()V
   #7 = Utf8               java/lang/System
   #8 = Class              #7             // java/lang/System
   #9 = Utf8               out
  #10 = Utf8               Ljava/io/PrintStream;
  #11 = NameAndType        #9:#10         // out:Ljava/io/PrintStream;
  #12 = Fieldref           #8.#11         // java/lang/System.out:Ljava/io/PrintStream;
  #13 = Utf8               Hello, World!
  #14 = String             #13            // Hello, World!
  #15 = Utf8               java/io/PrintStream
  #16 = Class              #15            // java/io/PrintStream
  #17 = Utf8               println
  #18 = Utf8               (Ljava/lang/String;)V
  #19 = NameAndType        #17:#18        // println:(Ljava/lang/String;)V
  #20 = Methodref          #16.#19        // java/io/PrintStream.println:(Ljava/lang/String;)V
  #21 = Utf8               HelloWorld
  #22 = Class              #21            // HelloWorld
  #23 = Utf8               Code
  #24 = Utf8               main
  #25 = Utf8               ([Ljava/lang/String;)V
  #26 = Utf8               globalValue
  #27 = Utf8               I
  #28 = Utf8               globalString
  #29 = Utf8               Ljava/lang/String;
{
  public int globalValue;
    descriptor: I
    flags: (0x0001) ACC_PUBLIC

  public java.lang.String globalString;
    descriptor: Ljava/lang/String;
    flags: (0x0001) ACC_PUBLIC

  public HelloWorld();
    descriptor: ()V
    flags: (0x0001) ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #6                  // Method java/lang/Object."<init>":()V
         4: return

  public static void main(java.lang.String...);
    descriptor: ([Ljava/lang/String;)V
    flags: (0x0089) ACC_PUBLIC, ACC_STATIC, ACC_VARARGS
    Code:
      stack=2, locals=1, args_size=1
         0: bipush        -3
         2: istore_1
         3: iload_1
         4: iconst_2
         5: if_icmpge     16
         8: getstatic     #12                 // Field java/lang/System.out:Ljava/io/PrintStream;
        11: ldc           #14                 // String Hello, World!
        13: invokevirtual #20                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
        16: return
}
```