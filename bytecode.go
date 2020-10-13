package bytecode

import (
    "fmt"
)

type Buffer struct {
    buffer []byte
}

func (b *Buffer) PushBytes(bytes ...byte) {
    b.buffer = append(b.buffer, bytes...)
}

func (b *Buffer) PushUInt16(values ...uint16) {
    for _, value := range values {
        b.buffer = append(b.buffer, Int16ToBinary(value)...)
    }
}

func (b *Buffer) PushUInt32(values ...uint32) {
    for _, value := range values {
        b.buffer = append(b.buffer, Int32ToBinary(value)...)
    }
}

func (b *Buffer) PushUInt64(values ...uint64) {
    for _, value := range values {
        b.buffer = append(b.buffer, Int64ToBinary(value)...)
    }
}

func (b *Buffer) PushFloat32(values ...float32) {
    for _, value := range values {
        b.buffer = append(b.buffer, Float32ToBinary(value)...)
    }
}

func (b *Buffer) PushFloat64(values ...float64) {
    for _, value := range values {
        b.buffer = append(b.buffer, Float64ToBinary(value)...)
    }
}

func (b *Buffer) Get() []byte {
    return b.buffer
}

type Constant struct {
    Tag  byte
    Info []byte
}

type ClassVisitor struct {
    MinorVersion   uint16
    MajorVersion   uint16
    AccessFlags    uint16
    Name           string
    SuperClassName string

    ConstantPoolSize  uint16
    ConstantPool      []Constant
    ConstantPoolCache map[string]uint16

    InterfaceLength uint16
    FieldLength     uint16
    MethodLength    uint16
    Methods         []*MethodVisitor
    AttributeLength uint16
}

func NewClass(majorVersion uint16, name, superClass string, accessFlags uint16) *ClassVisitor {
    visitor := ClassVisitor{
        MinorVersion:   0,
        MajorVersion:   majorVersion,
        AccessFlags:    accessFlags,
        Name:           name,
        SuperClassName: superClass,

        ConstantPoolSize:  1,
        ConstantPoolCache: make(map[string]uint16),
    }

    return &visitor
}

func (v *ClassVisitor) PushStringConstant(constant string) uint16 {
    position := v.PushUtf8Constant(constant)
    return v.PushUInt16Constant(String, position)
}

func (v *ClassVisitor) PushClassConstant(constant string) uint16 {
    position := v.PushUtf8Constant(constant)
    return v.PushUInt16Constant(Class, position)
}

func (v *ClassVisitor) PushUtf8Constant(constant string) uint16 {
    var bytes []byte
    bytes = append(bytes, Int16ToBinary(uint16(len(constant)))...)
    bytes = append(bytes, []byte(constant)...)
    return v.PushConstant(Utf8, bytes...)
}

func (v *ClassVisitor) PushUInt16Constant(tag byte, constants ...uint16) uint16 {
    var bytes []byte
    for _, constant := range constants {
        bytes = append(bytes, Int16ToBinary(constant)...)
    }
    return v.PushConstant(tag, bytes...)
}

func (v *ClassVisitor) PushUInt32Constant(tag byte, constants ...uint32) uint16 {
    var bytes []byte
    for _, constant := range constants {
        bytes = append(bytes, Int32ToBinary(constant)...)
    }
    return v.PushConstant(tag, bytes...)
}

func (v *ClassVisitor) PushUInt64Constant(tag byte, constants ...uint64) uint16 {
    var bytes []byte
    for _, constant := range constants {
        bytes = append(bytes, Int64ToBinary(constant)...)
    }
    return v.PushConstant(tag, bytes...)
}

func (v *ClassVisitor) PushFloat32Constant(tag byte, constants ...float32) uint16 {
    var bytes []byte
    for _, constant := range constants {
        bytes = append(bytes, Float32ToBinary(constant)...)
    }
    return v.PushConstant(tag, bytes...)
}

func (v *ClassVisitor) PushFloat64Constant(tag byte, constants ...float64) uint16 {
    var bytes []byte
    for _, constant := range constants {
        bytes = append(bytes, Float64ToBinary(constant)...)
    }
    return v.PushConstant(tag, bytes...)
}

func (v *ClassVisitor) PushNameAndTypeConstant(name, descriptor string) uint16 {
    namePosition := v.PushUtf8Constant(name)
    descriptorPosition := v.PushUtf8Constant(descriptor)

    return v.PushUInt16Constant(NameAndType, namePosition, descriptorPosition)
}

func (v *ClassVisitor) PushMethodRefConstant(class, name, descriptor string) uint16 {
    classPosition := v.PushClassConstant(class)
    natPosition := v.PushNameAndTypeConstant(name, descriptor)

    return v.PushUInt16Constant(Methodref, classPosition, natPosition)
}

func (v *ClassVisitor) PushFieldRefConstant(class, name, descriptor string) uint16 {
    classPosition := v.PushClassConstant(class)
    natPosition := v.PushNameAndTypeConstant(name, descriptor)

    return v.PushUInt16Constant(Fieldref, classPosition, natPosition)
}

func (v *ClassVisitor) PushConstant(tag byte, constant ...byte) uint16 {
    base16 := fmt.Sprintf("%x", constant)
    position, ok := v.ConstantPoolCache[base16]
    if ok {
        return position
    }

    data := Constant{
        Tag:  tag,
        Info: constant,
    }
    v.ConstantPool = append(v.ConstantPool, data)

    index := v.ConstantPoolSize

    v.ConstantPoolSize++

    v.ConstantPoolCache[base16] = index
    return index
}

func (v *ClassVisitor) AsBytecode() []byte {
    classPosition := v.PushClassConstant(v.Name)
    superClassPosition := v.PushClassConstant(v.SuperClassName)

    var methods [][]byte
    for _, method := range v.Methods {
        methods = append(methods, method.AsBytecode())
    }

    buffer := Buffer{}
    buffer.PushUInt32(0xcafebabe)
    buffer.PushUInt16(v.MinorVersion, v.MajorVersion)
    buffer.PushUInt16(v.ConstantPoolSize)

    for _, pool := range v.ConstantPool {
        buffer.PushBytes(pool.Tag)
        buffer.PushBytes(pool.Info...)
    }

    buffer.PushUInt16(v.AccessFlags)

    buffer.PushUInt16(classPosition)
    buffer.PushUInt16(superClassPosition)

    buffer.PushUInt16(v.InterfaceLength)
    buffer.PushUInt16(v.FieldLength)
    buffer.PushUInt16(v.MethodLength)
    for _, bytes := range methods {
        buffer.PushBytes(bytes...)
    }
    buffer.PushUInt16(v.AttributeLength)

    return buffer.Get()
}

func (v *ClassVisitor) NewMethod(accessFlags uint16, name, descriptor string) *MethodVisitor {
    visitor := MethodVisitor{
        Class:       v,
        AccessFlags: accessFlags,
        Name:        name,
        Descriptor:  descriptor,
    }
    v.MethodLength++
    v.Methods = append(v.Methods, &visitor)

    visitor.MaxLocals += DescriptorToStackSize(descriptor)[0]
    if accessFlags&AccStatic != AccStatic {
        visitor.MaxLocals++
    }
    return &visitor
}

type MethodVisitor struct {
    Class       *ClassVisitor
    AccessFlags uint16
    Name        string
    Descriptor  string

    MaxLocals uint16
    MaxStack  uint16

    StackObserver     uint16
    InvokeDescriptors []string

    Instructions    []Instruction
    InstructionData [][]byte
}

func (m *MethodVisitor) AddLdcInsn(javaType int, object interface{}) {
    var index uint16
    switch javaType {
    case TypeChar:
        fallthrough
    case TypeByte:
        fallthrough
    case TypeShort:
        fallthrough
    case TypeBoolean:
        fallthrough
    case TypeInt:
        index = m.Class.PushUInt32Constant(Integer, object.(uint32))
        m.AddInsn(Ldc, byte(index))
        break
    case TypeFloat:
        index = m.Class.PushFloat32Constant(Float, object.(float32))
        m.AddInsn(Ldc, byte(index))
        break
    case TypeDouble:
        index = m.Class.PushFloat64Constant(Double, object.(float64))
        m.AddInsn(Ldc2w, byte(index))
        break
    case TypeLong:
        index = m.Class.PushUInt64Constant(Long, object.(uint64))
        m.AddInsn(Ldc2w, byte(index))
        break
    case TypeString:
        index = m.Class.PushStringConstant(object.(string))
        m.AddInsn(Ldc, byte(index))
        break
    }
}

func (m *MethodVisitor) AddVarInsn(insn Instruction, args uint16) {
    m.AddInsn(insn, Int16ToBinary(args)...)
    m.MaxLocals++
}

func (m *MethodVisitor) AddMethodInsn(insn Instruction, instance, name, descriptor string) {
    methodRefPosition := m.Class.PushMethodRefConstant(instance, name, descriptor)

    m.InvokeDescriptors = append(m.InvokeDescriptors, descriptor)
    m.AddInsn(insn, Int16ToBinary(methodRefPosition)...)
}

func (m *MethodVisitor) AddFieldInsn(insn Instruction, instance, name, descriptor string) {
    fieldRefPosition := m.Class.PushFieldRefConstant(instance, name, descriptor)

    m.AddInsn(insn, Int16ToBinary(fieldRefPosition)...)
}

func (m *MethodVisitor) AddInsn(insn Instruction, data ...byte) {
    m.Instructions = append(m.Instructions, insn)
    m.InstructionData = append(m.InstructionData, data)
}

func (m *MethodVisitor) Maxs(maxStack, maxLocals uint16) {
    m.MaxStack = maxStack
    m.MaxLocals = maxLocals
}

func (m *MethodVisitor) AsBytecode() []byte {
    namePosition := m.Class.PushUtf8Constant(m.Name)
    descriptorPosition := m.Class.PushUtf8Constant(m.Descriptor)
    codePosition := m.Class.PushUtf8Constant("Code")

    insnBuffer := Buffer{}
    var stack uint16
    invokeIndex := 0
    for i, value := range m.Instructions {
        data := m.InstructionData[i]
        if value.StackIntakeFlag == FlagStackArgs {
            args := DescriptorToStackSize(m.InvokeDescriptors[invokeIndex])

            invokeIndex++
            stack -= args[0] + value.StackIntake
            if value.StackOutputFlag == FlagStackEmpty {
                stack = args[1] // empty the stack and add output size
            } else {
                stack += args[1] // add output size
            }
        } else if value.StackOutputFlag == FlagStackEmpty {
            stack = value.StackOutput
        } else {
            stack -= value.StackIntake
            stack += value.StackOutput
        }
        if stack > m.MaxStack {
            m.MaxStack = stack
        }
        insnBuffer.PushBytes(value.Opcode)
        insnBuffer.PushBytes(data...)
    }

    buffer := Buffer{}
    buffer.PushUInt16(m.AccessFlags)
    buffer.PushUInt16(namePosition)
    buffer.PushUInt16(descriptorPosition)
    buffer.PushUInt16(1) // attribute count

    // Code attribute
    buffer.PushUInt16(codePosition)
    buffer.PushUInt32(uint32(len(insnBuffer.Get()) + 12)) // code size
    buffer.PushUInt16(m.MaxStack, m.MaxLocals)
    buffer.PushUInt32(uint32(len(insnBuffer.Get()))) // insn size

    buffer.PushBytes(insnBuffer.Get()...) // instructions

    buffer.PushBytes(0, 0, 0, 0) // exceptions + attributes

    return buffer.Get()
}
