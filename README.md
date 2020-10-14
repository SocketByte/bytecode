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
        "my/package/HelloWorld", "java/lang/Object", []string{"java/lang/Cloneable"},
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

    // This also requires the Cloneable.clone method to be generated!
    clone := visitor.NewMethod(bytecode.AccPublic, "clone", "()Lmy/package/HelloWorld;")
    clone.AddTypeInsn(bytecode.New, "my/package/HelloWorld")
    clone.AddInsn(bytecode.Dup)
    clone.AddMethodInsn(bytecode.Invokespecial,
        "my/package/HelloWorld", "<init>", "()V")
    clone.AddInsn(bytecode.Areturn)

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
    main.AddLdcInsn("Hello, World!")
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
public class my.package.HelloWorld implements java.lang.Cloneable
  minor version: 0
  major version: 49
  flags: (0x0021) ACC_PUBLIC, ACC_SUPER
  this_class: #10                         // my/package/HelloWorld
  super_class: #4                         // java/lang/Object
  interfaces: 1, fields: 2, methods: 3, attributes: 1
Constant pool:
   #1 = Utf8               SourceFile
   #2 = Utf8               HelloWorld.java
   #3 = Utf8               java/lang/Object
   #4 = Class              #3             // java/lang/Object
   #5 = Utf8               <init>
   #6 = Utf8               ()V
   #7 = NameAndType        #5:#6          // "<init>":()V
   #8 = Methodref          #4.#7          // java/lang/Object."<init>":()V
   #9 = Utf8               my/package/HelloWorld
  #10 = Class              #9             // my/package/HelloWorld
  #11 = Methodref          #10.#7         // my/package/HelloWorld."<init>":()V
  #12 = Utf8               java/lang/System
  #13 = Class              #12            // java/lang/System
  #14 = Utf8               out
  #15 = Utf8               Ljava/io/PrintStream;
  #16 = NameAndType        #14:#15        // out:Ljava/io/PrintStream;
  #17 = Fieldref           #13.#16        // java/lang/System.out:Ljava/io/PrintStream;
  #18 = Utf8               Hello, World!
  #19 = String             #18            // Hello, World!
  #20 = Utf8               java/io/PrintStream
  #21 = Class              #20            // java/io/PrintStream
  #22 = Utf8               println
  #23 = Utf8               (Ljava/lang/String;)V
  #24 = NameAndType        #22:#23        // println:(Ljava/lang/String;)V
  #25 = Methodref          #21.#24        // java/io/PrintStream.println:(Ljava/lang/String;)V
  #26 = Utf8               java/lang/Cloneable
  #27 = Class              #26            // java/lang/Cloneable
  #28 = Utf8               Code
  #29 = Utf8               clone
  #30 = Utf8               ()Lmy/package/HelloWorld;
  #31 = Utf8               main
  #32 = Utf8               ([Ljava/lang/String;)V
  #33 = Utf8               globalValue
  #34 = Utf8               I
  #35 = Integer            43
  #36 = Utf8               ConstantValue
  #37 = Utf8               globalString
  #38 = Utf8               Ljava/lang/String;
{
  public final int globalValue;
    descriptor: I
    flags: (0x0011) ACC_PUBLIC, ACC_FINAL
    ConstantValue: int 43

  public java.lang.String globalString;
    descriptor: Ljava/lang/String;
    flags: (0x0001) ACC_PUBLIC

  public my.package.HelloWorld();
    descriptor: ()V
    flags: (0x0001) ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #8                  // Method java/lang/Object."<init>":()V
         4: return

  public my.package.HelloWorld clone();
    descriptor: ()Lmy/package/HelloWorld;
    flags: (0x0001) ACC_PUBLIC
    Code:
      stack=2, locals=1, args_size=1
         0: new           #10                 // class my/package/HelloWorld
         3: dup
         4: invokespecial #11                 // Method "<init>":()V
         7: areturn

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
         8: getstatic     #17                 // Field java/lang/System.out:Ljava/io/PrintStream;
        11: ldc           #19                 // String Hello, World!
        13: invokevirtual #25                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
        16: return
}
SourceFile: "HelloWorld.java"
```