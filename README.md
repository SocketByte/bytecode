# Bytecode
Very simple bytecode generation library for Go

## Basic example
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
```
Classfile HelloWorld.class
public class HelloWorld
  minor version: 0
  major version: 49
  flags: (0x0021) ACC_PUBLIC, ACC_SUPER
  this_class: #1                          // HelloWorld
  super_class: #3                         // java/lang/Object
  interfaces: 0, fields: 0, methods: 2, attributes: 0
Constant pool:
   #1 = Class              #2             // HelloWorld
   #2 = Utf8               HelloWorld
   #3 = Class              #4             // java/lang/Object
   #4 = Utf8               java/lang/Object
   #5 = Utf8               <init>
   #6 = Utf8               ()V
   #7 = Utf8               Code
   #8 = Class              #9             // java/lang/Object
   #9 = Utf8               java/lang/Object
  #10 = Utf8               <init>
  #11 = Utf8               ()V
  #12 = NameAndType        #10:#11        // "<init>":()V
  #13 = Methodref          #8.#12         // java/lang/Object."<init>":()V
  #14 = Utf8               main
  #15 = Utf8               ([Ljava/lang/String;)V
  #16 = Utf8               Code
  #17 = Class              #18            // java/lang/System
  #18 = Utf8               java/lang/System
  #19 = Utf8               out
  #20 = Utf8               Ljava/io/PrintStream;
  #21 = NameAndType        #19:#20        // out:Ljava/io/PrintStream;
  #22 = Fieldref           #17.#21        // java/lang/System.out:Ljava/io/PrintStream;
  #23 = String             #24            // Hello, World!
  #24 = Utf8               Hello, World!
  #25 = Class              #26            // java/io/PrintStream
  #26 = Utf8               java/io/PrintStream
  #27 = Utf8               println
  #28 = Utf8               (Ljava/lang/String;)V
  #29 = NameAndType        #27:#28        // println:(Ljava/lang/String;)V
  #30 = Methodref          #25.#29        // java/io/PrintStream.println:(Ljava/lang/String;)V
{
  public HelloWorld();
    descriptor: ()V
    flags: (0x0001) ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #13                 // Method java/lang/Object."<init>":()V
         4: return

  public static void main(java.lang.String...);
    descriptor: ([Ljava/lang/String;)V
    flags: (0x0089) ACC_PUBLIC, ACC_STATIC, ACC_VARARGS
    Code:
      stack=2, locals=1, args_size=1
         0: getstatic     #22                 // Field java/lang/System.out:Ljava/io/PrintStream;
         3: ldc           #23                 // String Hello, World!
         5: invokevirtual #30                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
         8: return
}
```