# Bytecode
Very simple bytecode generation library for Go

## Features
- Creating Java classes and methods
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

    bytecodeArray := visitor.AsBytecode()
    fmt.Println(bytecodeArray)
}
```
```
Classfile HelloWorld.class
public class HelloWorld
  minor version: 0
  major version: 49
  flags: (0x0021) ACC_PUBLIC, ACC_SUPER
  this_class: #22                         // HelloWorld
  super_class: #2                         // java/lang/Object
  interfaces: 0, fields: 0, methods: 2, attributes: 0
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
{
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
         0: getstatic     #12                 // Field java/lang/System.out:Ljava/io/PrintStream;
         3: ldc           #14                 // String Hello, World!
         5: invokevirtual #20                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
         8: return
}
```